const _exampleFile = Bun.file('input/example-a');
const _inputFile = Bun.file('input/input-a');

async function main() {
  const input = await _exampleFile.text();
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line

  for (const line of inputLines) {
    console.log(line);
  }
}

main();
