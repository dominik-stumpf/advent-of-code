export function solveProblem(input: string): number {
  const race = parseRace(input);
  let newDistanceRecordCounter = 0;

  for (let holdTime = 0; holdTime < race.timeBudget; holdTime += 1) {
    const dist = calcDistanceTraveled(race.timeBudget, holdTime);
    if (dist > race.distanceTraveled) {
      newDistanceRecordCounter += 1;
    }
  }

  return newDistanceRecordCounter;
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

function calcDistanceTraveled(timeBudget: number, holdTime: number) {
  return (timeBudget - holdTime) * holdTime;
}
