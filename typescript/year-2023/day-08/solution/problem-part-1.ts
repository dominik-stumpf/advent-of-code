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

function countTraversalSteps({ nodes, instructions }: Map): number {
  const originNode = 'AAA';
  const destinationNode = 'ZZZ';
  let stepCount = 0;
  let currentNode = originNode;
  let instructionIndex = 0;

  do {
    const instruction = instructions[instructionIndex];
    if (instruction === 'L') {
      currentNode = nodes[currentNode][0];
    } else {
      currentNode = nodes[currentNode][1];
    }

    instructionIndex = (instructionIndex + 1) % instructions.length;
    stepCount += 1;
  } while (currentNode !== destinationNode);

  return stepCount;
}
