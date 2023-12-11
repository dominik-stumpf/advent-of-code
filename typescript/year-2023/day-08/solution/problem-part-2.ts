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
  let stepCount = 0;
  let instructionIndex = 0;
  const nodeKeys = Object.keys(nodes);
  const originNodes = nodeKeys.filter((key) => key.at(-1) === 'A');
  const currentNodes = originNodes;
  const instructionIndices = instructions
    .split('')
    .map((instruction) => (instruction === 'L' ? 0 : 1));
  let time = performance.now();

  do {
    if (stepCount % 1e7 === 0) {
      const newTime = performance.now();
      console.log('step:', stepCount, 'elapsed:', newTime - time);
      time = newTime;
    }

    const instruction = instructionIndices[instructionIndex];
    for (let i = 0; i < currentNodes.length; i += 1) {
      currentNodes[i] = nodes[currentNodes[i]][instruction];
    }

    instructionIndex = (instructionIndex + 1) % instructionIndices.length;
    stepCount += 1;
  } while (!currentNodes.every((node) => node.at(-1) === 'Z'));

  return stepCount;
}
