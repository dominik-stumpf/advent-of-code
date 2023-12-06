export function solveProblem(input: string): number {
  const races = parseRaces(input);
  let result = 1;

  for (const race of races) {
    let newDistanceRecordCounter = 0;
    for (let i = 0; i < race.timeBudget; i += 1) {
      const dist = calcDistanceTraveled(race.timeBudget, i);
      if (dist > race.distanceTraveled) {
        newDistanceRecordCounter += 1;
      }
    }
    result *= newDistanceRecordCounter;
  }

  return result;
}

interface Race {
  timeBudget: number;
  distanceTraveled: number;
}

function parseRaces(rawRaces: string): Race[] {
  const [timeBudget, distanceTraveled] = rawRaces
    .split('\n')
    .slice(0, -1)
    .map((row) => row.split(':')[1].trim().split(/ +/).map(Number));

  return Array.from({ length: timeBudget.length }, (_, i) => ({
    distanceTraveled: distanceTraveled[i],
    timeBudget: timeBudget[i],
  }));
}

function calcDistanceTraveled(timeBudget: number, holdTime: number) {
  return (timeBudget - holdTime) * holdTime;
}
