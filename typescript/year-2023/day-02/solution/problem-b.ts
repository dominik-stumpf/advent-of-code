const _exampleFile = Bun.file('input/example-a');
const inputFile = Bun.file('input/input-a');

async function main() {
  const input = await inputFile.text();
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line
}

main();
