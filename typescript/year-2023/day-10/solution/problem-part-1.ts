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
  for (let x = 0; x < grid.length; x += 1) {
    const y = grid[x].findIndex((pipe) => pipe === 'S');
    if (y !== -1) {
      startPosition = { x, y };
    }
  }

  if (startPosition === undefined) {
    throw new Error('start position could not be found');
  }

  return { grid, startPosition };
}

const allDirections = ['up', 'right', 'down', 'left'] as const;

type Direction = typeof allDirections[number];

const pipeMap: Record<string, [Direction, Direction]> = {
  '|': ['up', 'down'],
  '-': ['left', 'right'],
  L: ['up', 'left'],
  J: ['up', 'right'],
  '7': ['left', 'down'],
  F: ['right', 'down'],
} as const;

function findValidAdjacentPipes(
  grid: PipeGrid['grid'],
  position: Vec2,
): Vec2[] {
  const currentDirections =
    pipeMap[grid[position.x][position.y]] ?? allDirections;

  return [
    { x: position.x + 1, y: position.y },
    { x: position.x - 1, y: position.y },
    { x: position.x, y: position.y + 1 },
    { x: position.x, y: position.y - 1 },
  ].filter((potentialPosition) => {
    if (potentialPosition.x < 0 || potentialPosition.y < 0) {
      return false;
    }

    const pipe = grid.at(potentialPosition.x)?.at(potentialPosition.y);
    const directions = pipe ? pipeMap[pipe] : undefined;
    if (directions === undefined) {
      return false;
    }

    for (const direction of directions) {
      for (const currentDirection of currentDirections) {
        if (checkIfPathable(currentDirection, direction)) {
          return true;
        }
      }
    }

    return false;
  });
}

function checkIfPathable(from: Direction, to: Direction): boolean {
  return (
    (from === 'left' && to === 'right') ||
    (from === 'right' && to === 'left') ||
    (from === 'up' && to === 'down') ||
    (from === 'down' && to === 'up')
  );
}

function traceLongestPath({ startPosition, grid }: PipeGrid): number {
  const tracedPositions: Vec2[] = [];
  let stepCount = 1;
  // const visualizer = traceVisualizer({ startPosition, grid });
  // console.log(startPosition);

  (function tracer(currentPosition: Vec2) {
    const adjacentPipes = findValidAdjacentPipes(grid, currentPosition);

    for (const { x, y } of adjacentPipes) {
      const isPipeAleadyTraced =
        tracedPositions.find((pos) => pos.x === x && pos.y === y) !== undefined;

      if (!isPipeAleadyTraced) {
        // setTimeout(() => {
        //   visualizer.updateCellColor({ x, y }, 'fgWhite');
        // }, stepCount * 250);

        tracedPositions.push({ x, y });
        stepCount += 1;
        tracer({ x, y });
      }
    }
  })(startPosition);

  return stepCount / 2;
}

function traceVisualizer({ grid }: PipeGrid) {
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
    buffer[x][y].color = colors[newColor];
    process.stdout.clearLine(1);
    process.stdout.cursorTo(0, 0);
    drawBufferToCanvas();
  }

  function drawBufferToCanvas() {
    process.stdout.write(
      buffer.reduce(
        (accumulator, row) =>
          `${accumulator +
          row.map(({ color, cell }) => color + cell + colors.reset).join('')
          }\n`,
        '',
      ),
    );
  }

  console.clear();

  return { updateCellColor };
}
