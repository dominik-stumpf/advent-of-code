export function solveProblem(input: string): number {
  const race = parseRace(input);
  const [result1, result2] = calcQuadraticEquation(
    1,
    -race.timeBudget,
    race.distanceTraveled,
  ).map(Math.floor);

  if (result1 === undefined && result2 === undefined) {
    throw new Error('invalid input value');
  }

  return result1 - result2;
}

interface Race {
  timeBudget: number;
  distanceTraveled: number;
}

function parseRace(rawRace: string): Race {
  const [timeBudget, distanceTraveled] = rawRace
    .split('\n')
    .slice(0, -1)
    .map((row) => Number(row.split(':')[1].trim().split(/ +/).join('')));

  return {
    distanceTraveled,
    timeBudget,
  };
}

type QuadraticEquationResult = [] | [number] | [number, number];

function calcQuadraticEquation(
  a: number,
  b: number,
  c: number,
): QuadraticEquationResult {
  let result: QuadraticEquationResult = [];
  const discriminant = b ** 2 - 4 * a * c;

  if (discriminant > 0) {
    const sqrtDiscriminant = Math.sqrt(discriminant);
    result = [
      ((-b + sqrtDiscriminant) / 2) * a,
      ((-b - sqrtDiscriminant) / 2) * a,
    ];
  } else if (discriminant === 0) {
    result = [(-b / 2) * a];
  }

  return result;
}
