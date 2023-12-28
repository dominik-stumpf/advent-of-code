import { PipeGrid, Vec2 } from "./problem-part-2";

export function traceVisualizer(grid: PipeGrid['grid']) {
  const colors = {
    reset: '\x1b[0m',
    bright: '\x1b[1m',
    dim: '\x1b[2m',
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
          `${accumulator +
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
