/**
 * OSOVM Opcode Benchmark
 * Runs sequential opcode executions and reports performance.
 */

const { execSync } = require('child_process');
const path = require('path');

const CLI_PATH = path.join(__dirname, '../src/cli.jl');

const OPCODES = [
  { opcode: "COUNCIL_APPROVE", args: { walletId: 0 } }, // Critical (Sui)
  { opcode: "PROJECT", args: { id: "bench-1", sector: "tech" } }, // Non-critical (Offloaded)
  { opcode: "STAKE", args: { amount: 10.0 } }, // Critical
  { opcode: "VEIL", args: { id: 1, f1: 0.98 } }, // Non-critical
  { opcode: "FINAL_SIGN", args: { walletId: 0 } }, // Critical
  { opcode: "IMPACT", args: { ase: 5.0 } }, // Critical
  { opcode: "BALANCE", args: { wallet: "shrine-01" } }, // Non-critical
  { opcode: "ORISA_OBATALA", args: {} }, // Non-critical
  { opcode: "TRANSFER", args: { to: "0x456", amount: 1.0 } }, // Critical
  { opcode: "EMIT", args: { event: "BenchmarkComplete" } } // Non-critical
];

function runBenchmark() {
  console.log("🚀 Starting OSOVM Opcode Benchmark...");
  
  const results = [];
  let totalTime = 0;

  for (let i = 0; i < OPCODES.length; i++) {
    const task = OPCODES[i];
    const taskJson = JSON.stringify(task);
    
    process.stdout.write(`[${i + 1}/${OPCODES.length}] Running ${task.opcode}... `);
    
    const start = Date.now();
    try {
      // Use absolute path for safety in exec
      const output = execSync(`julia /data/data/com.termux/files/home/osovm/src/cli.jl --task '${taskJson}'`, { encoding: 'utf8' });
      const elapsed = Date.now() - start;
      
      const data = JSON.parse(output);
      results.push({
        opcode: task.opcode,
        time: elapsed,
        status: data.status,
        offloaded: data.vm_result?.status === 'offloaded'
      });
      
      totalTime += elapsed;
      console.log(`${elapsed}ms ${results[results.length-1].offloaded ? '(OFFLOADED)' : ''}`);
    } catch (err) {
      console.log(`FAILED: ${err.message}`);
    }
  }

  console.log("\n" + "=".repeat(40));
  console.log("📊 BENCHMARK REPORT");
  console.log("=".repeat(40));
  console.log(`Total Time:   ${totalTime}ms`);
  console.log(`Avg per Op:   ${(totalTime / results.length).toFixed(2)}ms`);
  console.log("-".repeat(40));
  
  const sorted = [...results].sort((a, b) => b.time - a.time);
  console.log("Top Bottlenecks:");
  sorted.slice(0, 3).forEach((r, i) => {
    console.log(`${i + 1}. ${r.opcode}: ${r.time}ms`);
  });
  
  console.log("-".repeat(40));
  const offloadedCount = results.filter(r => r.offloaded).length;
  console.log(`Offloaded:    ${offloadedCount} / ${results.length}`);
  console.log("=".repeat(40));
}

runBenchmark();
