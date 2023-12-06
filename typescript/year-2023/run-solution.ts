import { readFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { program } from 'commander';
import { bench, run } from 'mitata';

function setupProgram() {
  program
    .name('solution-driver')
    .description('CLI to run any solution file and print result')
    .requiredOption(
      '-p, --part <number>',
      'part of the problem ranging from 1..2',
    )
    .requiredOption('-d, --day <number>', 'day of calendar ranging from 1..25')
    .option('--example', 'run solver using example file')
    .option('-b, --benchmark', 'measure performance using microbenchmark');

  return program.parse();
}

async function main() {
  const program = setupProgram();

  const options = program.opts();
  const day = options.day.padStart(2, '0');
  const { part, example, benchmark } = options;
  const solverPath = `@/day-${day}/solution/problem-part-${part}`;
  const inputPath = resolve(
    `day-${day}`,
    'input',
    `${example ? 'example' : 'input'}-part-${part}`,
  );
  const solution = await import(solverPath);
  const input = readFileSync(inputPath, 'utf-8');

  if (benchmark) {
    bench(solverPath, () => solution.solveProblem(input));
    await run({
      percentiles: false,
    });
  } else {
    console.log('solver result:', solution.solveProblem(input));
  }
}

main();
