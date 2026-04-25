#!/usr/bin/env bun
const URL = process.env.WORKERD_URL ?? "http://127.0.0.1:8787";
const TOTAL = Number(process.env.SOAK_TOTAL ?? 1000);
const PARALLEL = Number(process.env.SOAK_PAR ?? 10);
const DURATION_MS = Number(process.env.SOAK_DURATION_MS ?? 600_000);

const start = Date.now();
let claimed = 0; // requests handed out (no overshoot)
let succeeded = 0;
let failed = 0;
let errored = 0;

async function worker() {
  while (true) {
    if (Date.now() - start >= DURATION_MS) return;
    const myIdx = claimed;
    if (myIdx >= TOTAL) return;
    claimed++;
    try {
      const r = await fetch(URL + "/run", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          mode: "search",
          code: "async () => Object.keys(spec.paths).length",
        }),
      });
      const body = await r.json();
      if (body.status === "done") succeeded++;
      else failed++;
    } catch {
      errored++;
    }
  }
}

await Promise.all(Array.from({ length: PARALLEL }, worker));

const total = succeeded + failed + errored;
console.log(
  `claimed=${claimed} total=${total} succeeded=${succeeded} failed=${failed} errored=${errored} elapsed=${Date.now() - start}ms`,
);
if (failed > 0 || errored > 0) process.exit(1);
