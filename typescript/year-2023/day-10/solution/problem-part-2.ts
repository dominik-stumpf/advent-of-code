export function solveProblem(input: string): number {
  const pipeGrid = parsePipeGrid(input);
  return traceLongestPath(pipeGrid);
}

interface Vec2 {
  x: number;
  y: number;
}

interface PipeGrid {
  startPosition: Vec2;
  grid: string[][];
}

function parsePipeGrid(rawPipeGrid: string): PipeGrid {
  const grid = rawPipeGrid
    .trim()
    .split('\n')
    .map((row) => row.split(''));
  let startPosition: Vec2 | undefined;
  for (let y = 0; y < grid.length; y += 1) {
    const x = grid[y].findIndex((pipe) => pipe === 'S');
    if (x !== -1) {
      startPosition = { x, y };
    }
  }

  if (startPosition === undefined) {
    throw new Error('start position could not be found');
  }

  return { grid, startPosition };
}

type Direction = 'up' | 'right' | 'down' | 'left';

const pipeMap: Record<string, Direction[]> = {
  '|': ['up', 'down'],
  '-': ['left', 'right'],
  L: ['up', 'right'],
  J: ['up', 'left'],
  '7': ['left', 'down'],
  F: ['right', 'down'],
  S: ['right', 'down', 'up', 'left'],
} as const;

interface AdjacentPipe extends Vec2 {
  direction: Direction;
}

const oppositeDirection: Record<Direction, Direction> = {
  left: 'right',
  right: 'left',
  up: 'down',
  down: 'up',
} as const;

function findValidAdjacentPipes(grid: PipeGrid['grid'], position: Vec2) {
  const currentPipe = grid[position.y][position.x];
  const currentDirections = pipeMap[currentPipe];

  let adjacentPipes: AdjacentPipe[] = [
    { x: position.x + 1, y: position.y, direction: 'right' },
    { x: position.x - 1, y: position.y, direction: 'left' },
    { x: position.x, y: position.y + 1, direction: 'down' },
    { x: position.x, y: position.y - 1, direction: 'up' },
  ];

  adjacentPipes = adjacentPipes.filter(({ x, y, direction }) => {
    if (x < 0 || y < 0) {
      return false;
    }
    const potentialPipe = grid.at(y)?.at(x);
    const potentialDirections = potentialPipe
      ? pipeMap[potentialPipe]
      : undefined;
    if (potentialDirections === undefined) {
      return false;
    }

    return (
      currentDirections.includes(direction) &&
      potentialDirections.includes(oppositeDirection[direction])
    );
  });

  return adjacentPipes.map(({ x, y }) => ({ x, y }));
}

function traceLongestPath({ startPosition, grid }: PipeGrid): number {
  const tracedPositions: Vec2[] = [startPosition];
  let stepCount = 1;

  (function tracer(currentPosition: Vec2) {
    const adjacentPipes = findValidAdjacentPipes(grid, currentPosition);

    for (const { x, y } of adjacentPipes) {
      const isPipeAleadyTraced =
        tracedPositions.find((pos) => pos.x === x && pos.y === y) !== undefined;

      if (!isPipeAleadyTraced) {
        tracedPositions.push({ x, y });
        stepCount += 1;
        tracer({ x, y });
      }
    }
  })(startPosition);

  return stepCount / 2;
}
