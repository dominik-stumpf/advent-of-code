const _exampleFile = Bun.file('input/example-b');
const inputFile = Bun.file('input/input-b');

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

function getSmallestProcedureResult({ seeds, maps }: Almanac): number {
  let smallestProcedure = undefined;

  for (let i = 0; i < seeds.length; i += 2) {
    const seedRangeStart = seeds[i];
    const seedRange = seeds[i + 1];

    for (
      let seed = seedRangeStart;
      seed < seedRangeStart + seedRange;
      seed += 1
    ) {
      let seedProcedure = seed;
      for (const map of maps) {
        const newSeed = mapSeed(seedProcedure, map);
        seedProcedure = newSeed;
      }
      if (
        smallestProcedure === undefined ||
        seedProcedure < smallestProcedure
      ) {
        smallestProcedure = seedProcedure;
      }
    }
  }

  if (smallestProcedure === undefined) {
    throw new Error('no seed were calculated');
  }

  return smallestProcedure;
}

async function main() {
  const input = await inputFile.text();
  const almanac = parseAlmanac(input);
  const smallestProcedureResult = getSmallestProcedureResult(almanac);

  console.log('smallestProcedureResult:', smallestProcedureResult);
}

main();
