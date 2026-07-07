/**
 * Count car fleets by sorting by position and using a stack of arrival times.
 * Time: O(n log n), Space: O(n)
 */
export function solveCarFleet(target: number, position: number[], speed: number[]): number {
  const n = position.length;
  if (n === 0) {
    return 0;
  }
  const cars = position
    .map((pos, i) => ({ pos, speed: speed[i] }))
    .sort((a, b) => b.pos - a.pos);

  let fleets = 0;
  let lastTime = 0;
  for (const c of cars) {
    const time = (target - c.pos) / c.speed;
    if (time > lastTime) {
      fleets++;
      lastTime = time;
    }
  }
  return fleets;
}
