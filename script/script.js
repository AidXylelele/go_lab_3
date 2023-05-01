const greenFrameBtn = document.querySelector(".gf-script");
const diagonalMoveBtn = document.querySelector(".dm-script");
const windowLength = 800;
const halfFigure = 150;
const startPos = halfFigure;
const finishPos = windowLength - halfFigure;
const maxDistance = finishPos - startPos;
const step = 100;
const second = 1000;
const tenSeconds = 10000;
let position = startPos;
let distance = 0;

const config = {
  greenBtn: {
    green: null,
    bgrect: [200, 200, 600, 600],
    update: null,
  },
  diagonal: {
    figure: [startPos, startPos],
    update: null,
  },
  moveRight: {
    move: [0, 0],
  },
  moveLeft: {
    move: [0, 0],
  },
};

class UrlService {
  static url = "http://127.0.0.1:17000/?cmd=";

  static create = (options) => {
    let newUrl = this.url;
    for (const action in options) {
      newUrl += `${action}`;
      const arrayOfArgs = options[action];
      if (arrayOfArgs) newUrl += " " + arrayOfArgs.join(" ");
      newUrl += ",";
    }
    return newUrl.slice(0, -1);
  };
}

const moveFigure = async () => {
  await sendHTTPRequest(UrlService.create(config.diagonal));

  let flag = true;

  setTimeout(() => {
    flag = false;
  }, tenSeconds);

  while (flag) {
    while (position > startPos) {
      if (distance - step < 0) {
        config.moveRight.move = [-distance, -distance];
        sendHTTPRequest(UrlService.create(config.moveRight));
        position = startPos;
        distance = 0;
      } else {
        config.moveRight.move = [-step, -step];
        sendHTTPRequest(UrlService.create(config.moveRight));
        position -= step;
        distance -= step;
      }
      sendHTTPRequest(UrlService.create({ update: null }));
      await new Promise((resolve) => setTimeout(resolve, second));
    }

    while (position < finishPos) {
      if (distance + step > maxDistance) {
        config.moveLeft.move = [maxDistance - distance, maxDistance - distance];
        sendHTTPRequest(UrlService.create(config.moveLeft));
        position = finishPos;
        distance = maxDistance;
      } else {
        config.moveLeft.move = [step, step];
        sendHTTPRequest(UrlService.create(config.moveLeft));
        position += step;
        distance += step;
      }
      sendHTTPRequest(UrlService.create({ update: null }));
      await new Promise((resolve) => setTimeout(resolve, second));
    }
  }
};

const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

const sendHTTPRequest = async (url) => {
  try {
    const response = await fetch(url);
    if (response.ok) return response.text();
    throw new Error(`Request failed with status ${response.status}`);
  } catch (error) {
    console.error(error);
  }
};

greenFrameBtn.addEventListener("click", () => {
  const url = UrlService.create(config.greenBtn);

  sendHTTPRequest(url)
    .then((response) => console.log(response))
    .catch((error) => console.error(error));
});

diagonalMoveBtn.addEventListener("click", async () => {
  const url = UrlService.create(config.diagonal);
  await sendHTTPRequest(url);
  await moveFigure();
});
