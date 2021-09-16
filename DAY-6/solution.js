const pushNews = (news, startTime) =>
  new Promise((resolve) =>
    setTimeout(() => {
      console.log(`${news} cost ${Date.now() - startTime}`);
      resolve(Date.now());
    }, 3000)
  );

const start = Date.now();
Promise.all(
  ["中秋節來了", "記得", "不要戶外烤肉～"].map((news) => pushNews(news, start))
).then((allNews) =>
  allNews.map((finishTimes, index) =>
    console.log(`"news ${index} is sent at ${finishTimes}"`)
  )
);
