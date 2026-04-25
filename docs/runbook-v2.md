# CreateOS MCP v2 Runbook

## Components

- **mcp container**: Go binary serving MCP tools.
- **workerd container**: V8 sandbox sidecar at `127.0.0.1:8787`.
- Comms over loopback. workerd never reachable from outside the pod.

## Health

- `mcp /readyz` (port 8080) → 200 only when workerd `/health` responds 200.
- `workerd /health` returns `{ok, specVersion, endpointCount, jobsRunning, jobStoreSize}`.
- `mcp /metrics` exposes Prom: `codemode_run_duration_seconds`, `codemode_jobs_running`, `codemode_jobstore_size`, `codemode_api_call_duration_seconds`.

## Common failures

| Symptom | Likely cause | Fix |
|---|---|---|
| `/readyz` 503 indefinitely | workerd cannot reach backend `/api-docs/openapi.yaml` | Check `BACKEND_URL` secret + backend uptime |
| All `execute` returns `errorKind:infra` | workerd dispatcher crashed | Check workerd container logs; restart pod |
| Spec stale | backend shipped new endpoint, sidecar cached old spec | Rolling restart |
| `jobMissing` for valid id | Job evicted (TTL 30 min after terminal) | Re-run execute |
| Many `errorKind:timeout` | Long-running ops hitting 600s wall cap | Investigate why ops take so long; consider backend optimisation |

## Rollback

```bash
kubectl set image pod/createos-mcp mcp=createos-mcp:v1.x.y-final
kubectl delete pod -l app=createos-mcp,version=v2
```

V1 image has no workerd dependency; pod will become ready immediately.

## Capacity

- Each isolate: ~10 MB RSS peak.
- jobStore: bounded by `JOB_TTL_MS` (30 min) × concurrency.
- workerd binary: ~80 MB image footprint.

## Logs

- `mcp` logs: tool name, duration, jobId, workerd status.
- `workerd dispatcher` logs: jobId, mode, duration, outcome (no model code body, no auth).
- `workerd api-proxy` logs: method, path, status, duration (no headers, no body).
