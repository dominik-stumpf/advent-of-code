interface GameRecord {
  id: number;
  setOfCubes: {
    red: number;
    green: number;
    blue: number;
  }[];
}

const initialSetOfCubes: GameRecord['setOfCubes'][number] = {
  red: 12,
  green: 13,
  blue: 14,
};

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

function validateGameRecord(record: GameRecord): boolean {
  return record.setOfCubes.every(
    (setOfCube) =>
      setOfCube.red <= initialSetOfCubes.red &&
      setOfCube.green <= initialSetOfCubes.green &&
      setOfCube.blue <= initialSetOfCubes.blue,
  );
}


export function solveProblem(input: string): number {
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line

  const sum = inputLines.reduce((accumulator, line) => {
    const gameRecord = parseGameRecord(line);
    const isGameRecordValid = validateGameRecord(gameRecord);

    return accumulator + (isGameRecordValid ? gameRecord.id : 0);
  }, 0);

  return sum;
}
