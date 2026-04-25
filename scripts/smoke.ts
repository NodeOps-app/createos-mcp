#!/usr/bin/env bun
const URL = process.env.WORKERD_URL ?? "http://127.0.0.1:8787";
const N = Number(process.env.SMOKE_N ?? 50);

const start = Date.now();
const results = await Promise.all(
  Array.from({ length: N }, async (_, i) => {
    const t = performance.now();
    const r = await fetch(URL + "/run", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify({
        mode: "search",
        code: "async () => Object.keys(spec.paths).length",
      }),
    });
    const body = await r.json();
    return { i, dur: performance.now() - t, ok: body.status === "done" };
  }),
);
const ok = results.filter((r) => r.ok).length;
const durs = results.map((r) => r.dur).sort((a, b) => a - b);
const p = (q: number) => durs[Math.floor(q * durs.length)];
console.log(`ok=${ok}/${N} elapsed=${Date.now() - start}ms p50=${p(0.5).toFixed(0)}ms p95=${p(0.95).toFixed(0)}ms`);
if (ok < N) process.exit(1);
