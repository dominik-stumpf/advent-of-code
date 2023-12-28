import { traceVisualizer } from './trace-visualizer';

let visualizer: undefined | ReturnType<typeof traceVisualizer>;

export function solveProblem(input: string): number {
  const pipeGrid = parsePipeGrid(input);
  const { potentialEnclosedTiles, tracedPipeGrid } = tracePipeGrid(pipeGrid);

  visualizer = traceVisualizer(tracedPipeGrid);
  visualizeGrid(tracedPipeGrid, visualizer);

  const enclosedTiles = filterPotentialEnclosedTiles({ potentialEnclosedTiles, tracedPipeGrid });

  for (const enclosedTile of enclosedTiles) {
    visualizer.updateCellColor({ x: enclosedTile.x, y: enclosedTile.y }, 'fgGreen');
  }

  return enclosedTiles.length;
}

export interface Vec2 {
  x: number;
  y: number;
}

export interface PipeGrid {
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

function findPotentialAdjacentPipes(grid: PipeGrid['grid'], position: Vec2): AdjacentPipe[] {
  const adjacentPipes: AdjacentPipe[] = [
    { x: position.x + 1, y: position.y, direction: 'right' },
    { x: position.x - 1, y: position.y, direction: 'left' },
    { x: position.x, y: position.y + 1, direction: 'down' },
    { x: position.x, y: position.y - 1, direction: 'up' },
  ];

  return adjacentPipes.filter(({ x, y }) => {
    if (x < 0 || y < 0) {
      return false;
    }
    const potentialPipe = grid.at(y)?.at(x);

    return potentialPipe !== undefined;
  }) as AdjacentPipe[]
}

function findValidAdjacentPipes(grid: PipeGrid['grid'], position: Vec2): Vec2[] {
  const currentPipe = grid[position.y][position.x];
  const currentDirections = pipeMap[currentPipe];
  const potentialAdjacentPipes = findPotentialAdjacentPipes(grid, position);
  const adjacentPipes = potentialAdjacentPipes.filter(({ x, y, direction }) => {
    const potentialPipe = grid[y][x];
    // if (potentialPipe === '!') {
    //   return true;
    // }
    const potentialDirections = pipeMap[potentialPipe]
    if (potentialDirections === undefined) {
      return false;
    }

    // if (currentDirections === undefined) {
    //   return true;
    // }

    return (
      currentDirections.includes(direction) &&
      potentialDirections.includes(oppositeDirection[direction])
    );
  });

  return adjacentPipes.map(({ x, y }) => ({ x, y }));
}

function tracePipeGrid({ startPosition, grid }: PipeGrid) {
  const tracedPositions: Vec2[] = [startPosition];
  const tracedPipeGrid = Array.from({ length: grid.length }, () =>
    Array(grid[0].length).fill('!'),
  ) as PipeGrid['grid'];

  (function tracer(currentPosition: Vec2) {
    const adjacentPipes = findValidAdjacentPipes(grid, currentPosition);

    for (const { x, y } of adjacentPipes) {
      const isPipeAleadyTraced =
        tracedPositions.find((pos) => pos.x === x && pos.y === y) !== undefined;

      if (!isPipeAleadyTraced) {
        tracedPositions.push({ x, y });
        tracedPipeGrid[y][x] = grid[y][x];
        tracer({ x, y });
      }
    }
  })(startPosition);

  const potentialEnclosedTiles: Vec2[] = [];

  for (let y = 0; y < tracedPipeGrid.length; y += 1) {
    const row = tracedPipeGrid[y];
    const firstIndex = row.findIndex((pipe) => pipe !== '!');
    const lastIndex = row.findLastIndex((pipe) => pipe !== '!');
    if (firstIndex !== -1 && lastIndex !== -1) {
      for (let x = firstIndex; x < lastIndex; x += 1) {
        if (tracedPipeGrid[y][x] === '!') {
          potentialEnclosedTiles.push({ x, y });
          tracedPipeGrid[y][x] = '?';
        }
      }
    }
  }

  return { tracedPipeGrid, potentialEnclosedTiles };
}

function filterPotentialEnclosedTiles({
  potentialEnclosedTiles,
  tracedPipeGrid: grid,
}: { potentialEnclosedTiles: Vec2[]; tracedPipeGrid: PipeGrid['grid'] }) {
  let firstPipePosition: Vec2 | undefined;

  top: for (let y = 0; y < grid.length; y += 1) {
    for (let x = 0; x < grid[0].length; x += 1) {
      const pipe = grid[y][x]
      if (pipe !== '!' && pipe !== '?') {
        firstPipePosition = { x, y }
        break top;
      }
    }
  }

  if (firstPipePosition === undefined) {
    throw new Error('no pipe found');
  }


  const tracedPositions: Vec2[] = [firstPipePosition];
  let step = 0;

  (function tracer(currentPosition: Vec2) {
    const adjacentPipes = findValidAdjacentPipes(grid, currentPosition);

    for (const { x, y } of adjacentPipes) {
      const isPipeAleadyTraced =
        tracedPositions.find((pos) => pos.x === x && pos.y === y) !== undefined;

      if (!isPipeAleadyTraced) {
        tracedPositions.push({ x, y });
        step += 1;
        setTimeout(() => {
          visualizer?.updateCellColor({ x, y }, 'fgCyan')
        }, step * 2);
        tracer({ x, y });
      }
    }
  })(firstPipePosition);


  return potentialEnclosedTiles;
  // return potentialEnclosedTiles.filter((tile) => {
  //   const tracedPositions: Vec2[] = [tile];
  //
  //   function tracer(currentPosition: Vec2) {
  //     const adjacentPipes = findValidAdjacentPipes(grid, currentPosition);
  //
  //     for (const { x, y } of adjacentPipes) {
  //       const isPipeAleadyTraced =
  //         tracedPositions.find((pos) => pos.x === x && pos.y === y) !==
  //         undefined;
  //
  //       if (!isPipeAleadyTraced) {
  //         if (grid[y][x] === '!') {
  //           return false;
  //         }
  //         tracedPositions.push({ x, y });
  //         return tracer({ x, y });
  //       }
  //     }
  //
  //     return true;
  //   };
  //
  //   const adjacentPipes = findPotentialAdjacentPipes(grid, tile);
  //   for (const { x, y } of adjacentPipes) {
  //     const pipe = grid[y][x];
  //     if (pipe === '!') {
  //       return false;
  //     }
  //     if (pipe === '?') {
  //       continue;
  //     }
  //     if (tracer({ x, y })) {
  //       return false;
  //     }
  //   }
  //   return true;
  // });
}

function visualizeGrid(grid: PipeGrid['grid'], visualizer: ReturnType<typeof traceVisualizer>) {
  for (let x = 0; x < grid.length; x += 1) {
    for (let y = 0; y < grid[0].length; y += 1) {
      const pipe = grid[y][x];
      if (pipe === '!') {
        visualizer.updateCellColor({ x, y }, 'fgRed');
      } else if (pipe === '?') {
        visualizer.updateCellColor({ x, y }, 'fgBlue');
      } else {
        visualizer.updateCellColor({ x, y }, 'fgWhite');
      }
    }
  }
}
