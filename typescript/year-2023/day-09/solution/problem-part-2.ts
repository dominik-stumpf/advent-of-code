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
  let prediction = 0;
  let nextRow = history;
  const lastValues: number[] = [history[0]];

  while (!nextRow.every((value, _, array) => array[0] === value)) {
    nextRow = createNextRow(nextRow);
    lastValues.push(nextRow[0]);
  }

  for (let i = lastValues.length - 1; i >= 0; i -= 1) {
    prediction = lastValues[i] - prediction;
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
