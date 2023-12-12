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
  // const visualizer = _traceVisualizer({ startPosition, grid });

  (function tracer(currentPosition: Vec2) {
    const adjacentPipes = findValidAdjacentPipes(grid, currentPosition);

    for (const { x, y } of adjacentPipes) {
      const isPipeAleadyTraced =
        tracedPositions.find((pos) => pos.x === x && pos.y === y) !== undefined;

      if (!isPipeAleadyTraced) {
        // setTimeout(() => {
        //   visualizer.updateCellColor({ x, y }, 'fgWhite');
        // }, stepCount);

        tracedPositions.push({ x, y });
        stepCount += 1;
        tracer({ x, y });
      }
    }
  })(startPosition);

  return stepCount / 2;
}

function _traceVisualizer({ grid }: PipeGrid) {
  const colors = {
    reset: '\x1b[0m',
    bright: '\x1b[1m',
    dim: '\x1b[2m',
    underscore: '\x1b[4m',
    blink: '\x1b[5m',
    reverse: '\x1b[7m',
    hidden: '\x1b[8m',
    fgBlack: '\x1b[30m',
    fgRed: '\x1b[31m',
    fgGreen: '\x1b[32m',
    fgYellow: '\x1b[33m',
    fgBlue: '\x1b[34m',
    fgMagenta: '\x1b[35m',
    fgCyan: '\x1b[36m',
    fgWhite: '\x1b[37m',
    fgGray: '\x1b[90m',
  };

  const buffer = Array.from({ length: grid.length }, (_, i) =>
    grid[i].map((cell) => {
      let color: string;
      switch (cell) {
        case 'S':
          color = colors.fgYellow;
          break;
        default:
          color = colors.fgBlack;
          break;
      }

      return { cell, color };
    }),
  );

  function updateCellColor({ x, y }: Vec2, newColor: keyof typeof colors) {
    buffer[y][x].color = colors[newColor];
    process.stdout.cursorTo(x, y);
    process.stdout.write(colors[newColor] + grid[y][x] + colors.reset);
  }

  function drawBufferToCanvas() {
    process.stdout.cursorTo(0, 0);
    process.stdout.write(
      buffer.reduce(
        (accumulator, row) =>
          `${
            accumulator +
            row.map(({ color, cell }) => color + cell + colors.reset).join('')
          } \n`,
        '',
      ),
    );
  }

  console.clear();
  if (
    process.stdout.columns >= grid.length &&
    process.stdout.rows >= grid[0].length
  ) {
    drawBufferToCanvas();
  }

  return { updateCellColor };
}
