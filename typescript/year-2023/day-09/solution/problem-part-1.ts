export function solveProblem(input: string): number {
  const histories = parseHistories(input);

  return histories.reduce(
    (accumulator, history) => accumulator + extrapolateHistory(history),
    0,
  );
}

function parseHistories(rawHistory: string): number[][] {
  return rawHistory
    .trim()
    .split('\n')
    .map((line) => line.split(' ').map(Number));
}

function extrapolateHistory(history: number[]): number {
  let prediction = history[history.length - 1];
  let nextRow = history;

  while (!nextRow.every((value, _, array) => array[0] === value)) {
    nextRow = createNextRow(nextRow);
    prediction += nextRow[nextRow.length - 1];
  }

  return prediction;
}

function createNextRow(currentRow: number[]): number[] {
  const nextRow: number[] = [];

  for (let i = 0; i < currentRow.length - 1; i += 1) {
    nextRow.push(currentRow[i + 1] - currentRow[i]);
  }

  return nextRow;
}
