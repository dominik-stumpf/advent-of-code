import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';

function main() {
  const solutionTemplate = readFileSync('solution-template.ts', 'utf-8');

  for (let dayIndex = 1; dayIndex <= 25; dayIndex += 1) {
    const day = dayIndex.toString().padStart(2, '0');

    mkdirSync(`day-${day}/solution/`, { recursive: true, mode: 0o755 });
    mkdirSync(`day-${day}/input/`, { recursive: true, mode: 0o755 });

    for (let part = 1; part <= 2; part += 1) {
      const problem = `day-${day}/solution/problem-part-${part}.ts`
      const input = `day-${day}/input/input-part-${part}`
      const example = `day-${day}/input/example-part-${part}`

      if (!existsSync(problem)) {
        writeFileSync(problem, solutionTemplate, {});
      }
      if (!existsSync(input)) {
        writeFileSync(input, '', {});
      }
      if (!existsSync(example)) {
        writeFileSync(example, '', {});
      }
    }
  }

  console.log('project structure is successfully created');
}

main();
