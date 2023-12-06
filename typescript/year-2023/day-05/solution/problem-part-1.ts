interface MapRange {
  destinationRangeStart: number;
  sourceRangeStart: number;
  range: number;
}

interface Almanac {
  seeds: number[];
  maps: MapRange[][];
}

function parseAlmanac(rawAlmanac: string): Almanac {
  const [rawSeeds, ...rawMaps] = rawAlmanac
    .split('\n\n')
    .map((section) => section.split(':')[1].trim());
  const seeds = rawSeeds.split(' ').map(Number);
  const maps = rawMaps.map((map) =>
    map
      .split('\n')
      .map((ranges) => ranges.split(' ').map(Number))
      .map(([destinationRangeStart, sourceRangeStart, range]) => ({
        sourceRangeStart,
        destinationRangeStart,
        range,
      })),
  );

  return { seeds, maps };
}

function mapSeed(seed: number, map: MapRange[]): number {
  for (const mapRange of map) {
    if (
      seed >= mapRange.sourceRangeStart &&
      seed <= mapRange.sourceRangeStart + mapRange.range
    ) {
      return (
        seed - (mapRange.sourceRangeStart - mapRange.destinationRangeStart)
      );
    }
  }

  return seed;
}

function mapAlmanacSeeds({ seeds, maps }: Almanac): number[] {
  const procedures: number[] = [];

  for (const seed of seeds) {
    let seedProcedure = seed;
    for (const map of maps) {
      const newSeed = mapSeed(seedProcedure, map);
      seedProcedure = newSeed;
    }
    procedures.push(seedProcedure);
  }

  return procedures;
}

export function solveProblem(input: string): number {
  const almanac = parseAlmanac(input);
  const lastProcedureResults = mapAlmanacSeeds(almanac);
  const smallestProcedureResult = Math.min(...lastProcedureResults);

  return smallestProcedureResult;
}
