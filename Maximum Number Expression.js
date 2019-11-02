function maxNumberExpression(numns, N, K) {
  sum = [0, ...numns];
  sum.reduce((pre, cur, index) => (sum[index] = pre + cur), 0);
  dp = new Array(N + 1).fill(0).map(_ => new Array(K + 1).fill(0));
  dp.map((line, index) => (line[0] = sum[index]));
  for (let i = 0; i <= N; i++) {
    dp[i][0] = sum[i];
  }
  for (let i = 2; i <= N; i++) {
    for (let j = 1; j <= K && j <= i - 1; j++) {
      for (let p = 2; p <= i; p++) {
        dp[i][j] = Math.max(dp[i][j], dp[p - 1][j - 1] * (sum[i] - sum[p - 1]));
      }
    }
  }

  return dp[N][K];
}
