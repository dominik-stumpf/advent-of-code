import { existsSync, readFileSync, unlinkSync, writeFileSync } from 'node:fs';

function main() {
  for (let dayIndex = 1; dayIndex <= 25; dayIndex += 1) {
    const day = dayIndex.toString().padStart(2, '0');

    for (let part = 1; part <= 2; part += 1) {
      const partLetter = 'ab'[part - 1];
      const oldInput = `day-${day}/input/input-${partLetter}`;
      const oldExample = `day-${day}/input/example-${partLetter}`;

      const input = `day-${day}/input/input-part-${part}`;
      const example = `day-${day}/input/example-part-${part}`;

      if (existsSync(oldInput)) {
        const oldFile = readFileSync(oldInput, 'utf-8');
        writeFileSync(input, oldFile, {});
        unlinkSync(oldInput);
      }
      if (existsSync(oldExample)) {
        const oldFile = readFileSync(oldExample, 'utf-8');
        writeFileSync(example, oldFile, {});
        unlinkSync(oldExample);
      }
    }
  }

  console.log('migration done');
}

main();
