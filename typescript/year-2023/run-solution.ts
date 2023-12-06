import { program } from 'commander';

function setupProgram() {
  program
    .name('solution-driver')
    .description('CLI to run any solution file and print result')
    .requiredOption(
      '-p, --part <number>',
      'part of the problem ranging from 1..2',
    )
    .requiredOption('-d, --day <number>', 'day of calendar ranging from 1..25')
    .option('--example', 'run solver using example file');

  return program.parse();
}

function main() {
  const program = setupProgram();

  const options = program.opts();
  const day = options.day.padStart(2, '0');
  const { part } = options;
  const solverPath = `@/day-${day}/solution/problem-part-${part}`;

  // const solution = await import(solverPath);
  console.log(options, solverPath);
}

main();
