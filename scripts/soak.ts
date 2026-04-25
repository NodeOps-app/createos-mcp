#!/usr/bin/env bun
const URL = process.env.WORKERD_URL ?? "http://127.0.0.1:8787";
const TOTAL = Number(process.env.SOAK_TOTAL ?? 1000);
const PARALLEL = Number(process.env.SOAK_PAR ?? 10);
const DURATION_MS = Number(process.env.SOAK_DURATION_MS ?? 600_000);

const start = Date.now();
let done = 0;
let failed = 0;

async function worker() {
  while (done < TOTAL && Date.now() - start < DURATION_MS) {
    const r = await fetch(URL + "/run", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify({
        mode: "search",
        code: "async () => Object.keys(spec.paths).length",
      }),
    });
    const body = await r.json();
    if (body.status !== "done") failed++;
    done++;
  }
}

await Promise.all(Array.from({ length: PARALLEL }, worker));

console.log(`done=${done} failed=${failed} elapsed=${Date.now() - start}ms`);
if (failed > 0) process.exit(1);
