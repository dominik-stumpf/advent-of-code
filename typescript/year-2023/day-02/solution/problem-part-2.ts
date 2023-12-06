interface GameRecord {
  id: number;
  setOfCubes: {
    red: number;
    green: number;
    blue: number;
  }[];
}

function parseGameRecord(recordLine: string): GameRecord {
  const matchResult = recordLine.match(/Game (\d+): (.+)/);

  if (matchResult === null || matchResult.length !== 3) {
    throw new Error(`invalid game record line found: ${recordLine}`);
  }

  const [, rawId, rawSetOfCubes] = matchResult;

  const id = Number(rawId);
  const setOfCubes = rawSetOfCubes.split(';').map((setOfCube) => {
    const cubes = setOfCube.trim().split(', ');
    const parsedCubes = cubes.map((cube) => {
      const [cubeCount, cubeColor] = cube.split(' ');
      return [cubeColor, Number(cubeCount)];
    });
    return { red: 0, green: 0, blue: 0, ...Object.fromEntries(parsedCubes) };
  });

  return { id, setOfCubes };
}

function calcGameRecordPower({ setOfCubes }: GameRecord): number {
  const tracker = { red: 0, green: 0, blue: 0 };
  for (let i = 0; i < setOfCubes.length; i += 1) {
    const setOfCube = setOfCubes[i];
    for (const stringColor in tracker) {
      const color = stringColor as keyof typeof tracker;
      if (tracker[color] < setOfCube[color]) {
        tracker[color] = setOfCube[color];
      }
    }
  }

  return Object.entries(tracker).reduce(
    (accumulator, current) => accumulator * current[1],
    1,
  );
}

export function solveProblem(input: string): number {
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line

  const sum = inputLines.reduce((accumulator, line) => {
    const gameRecord = parseGameRecord(line);
    const gameRecordPower = calcGameRecordPower(gameRecord);

    return accumulator + gameRecordPower;
  }, 0);

  return sum;
}
