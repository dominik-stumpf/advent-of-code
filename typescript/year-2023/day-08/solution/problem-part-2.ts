export function solveProblem(input: string): number {
  const map = parseMap(input);
  const steps = countTraversalSteps(map);

  return steps;
}

interface Map {
  instructions: string;
  nodes: Record<string, [string, string]>;
}

function parseMap(rawMap: string): Map {
  const [instructions, rawNodes] = rawMap.trim().split('\n\n');
  const nodes = Object.fromEntries(
    rawNodes
      .split('\n')
      .map((node) => node.split(' = '))
      .map(([currentNode, next]) => [
        currentNode,
        next.slice(1, -1).split(', '),
      ]),
  );

  return { instructions, nodes: nodes as Map['nodes'] };
}

function countTraversalSteps(map: Map): number {
  const originNodes = Object.keys(map.nodes).filter(
    (key) => key.at(-1) === 'A',
  );
  const exhaustLimits: number[] = [];

  for (const originNode of originNodes) {
    const traverser = traverseThread(map, originNode);
    let stepCounter = 0;
    for (const node of traverser) {
      stepCounter += 1;
      if (node.at(-1) === 'Z') {
        exhaustLimits.push(stepCounter);
        break;
      }
    }
  }

  return calcLeastCommonMultiple(exhaustLimits);
}

function* traverseThread({ nodes, instructions }: Map, startingNode: string) {
  let instructionIndex = 0;
  let currentNode = startingNode;
  const instructionIndices = instructions
    .split('')
    .map((instruction) => (instruction === 'L' ? 0 : 1));

  while (true) {
    const instruction = instructionIndices[instructionIndex];
    currentNode = nodes[currentNode][instruction];
    yield currentNode;
    instructionIndex = (instructionIndex + 1) % instructionIndices.length;
  }
}

function calcLeastCommonMultiple(numbers: number[]): number {
  if (numbers.length < 2) {
    throw new Error(
      'at least two numbers are required to find the least common multiple.',
    );
  }
  function gcd(a: number, b: number): number {
    return b === 0 ? a : gcd(b, a % b);
  }
  function lcm(a: number, b: number): number {
    return (a * b) / gcd(a, b);
  }

  let result = numbers[0];
  for (let i = 1; i < numbers.length; i += 1) {
    result = lcm(result, numbers[i]);
  }

  return result;
}
