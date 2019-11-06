class Activity {
  start: number;
  end: number;
  weight: number;

  constructor(start: number, end: number, weight: number) {
    this.start = start;
    this.end = end;
    this.weight = weight;
  }
}

function getBestSchedule(activities: Activity[]): number {
  let res = new Array(activities.length + 1)
    .fill(0)
    .map((_, index) => (index == 0 ? 0 : activities[index - 1].weight));

  activities.sort((a, b) => a.end - b.end);
  for (let i = 1; i <= activities.length; i++) {
    for (let j = i - 1; j >= 0; j--) {
      let act = activities[j];
      if (act.start > activities[j].end) {
        res[i] =
          res[i - 1] > res[j] + activities[i].weight
            ? res[i - 1]
            : res[j] + activities[i].weight;
      }
    }
  }
  console.log(res);

  return res[activities.length];
}

function test() {
  let acts = [
    new Activity(1, 6, 3),
    new Activity(2, 3, 2),
    new Activity(28, 49, 21),
    new Activity(1, 19, 28)
  ];

  console.log(getBestSchedule(acts));
}

test();
