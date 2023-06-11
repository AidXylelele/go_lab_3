import { sendHTTPRequest } from "./request-api.js";
import { UrlService } from "./services.js";

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

const greenBtn = {
  green: null,
  bgrect: [200, 200, 600, 600],
  update: null,
};
const diagonal = {
  figure: [startPos, startPos],
  update: null,
};
const moveRight = {
  move: [0, 0],
};
const moveLeft = {
  move: [0, 0],
};

const moveFigure = async () => {
  await sendHTTPRequest(UrlService.create(diagonal));

  let flag = true;

  setTimeout(() => {
    flag = false;
  }, tenSeconds);

  for (let i = 0; flag && i < 1000; i++) {
    while (position > startPos) {
      if (distance - step < 0) {
        moveRight.move = [-distance, -distance];
        await sendHTTPRequest(UrlService.create(moveRight));
        position = startPos;
        distance = 0;
      } else {
        moveRight.move = [-step, -step];
        await sendHTTPRequest(UrlService.create(moveRight));
        position -= step;
        distance -= step;
      }
      await sendHTTPRequest(UrlService.create({ update: null }));
      await new Promise((resolve) => setTimeout(resolve, second));
    }

    while (position < finishPos) {
      if (distance + step > maxDistance) {
        moveLeft.move = [maxDistance - distance, maxDistance - distance];
        await sendHTTPRequest(UrlService.create(moveLeft));
        position = finishPos;
        distance = maxDistance;
      } else {
        moveLeft.move = [step, step];
        await sendHTTPRequest(UrlService.create(moveLeft));
        position += step;
        distance += step;
      }
      await sendHTTPRequest(UrlService.create({ update: null }));
      await new Promise((resolve) => setTimeout(resolve, second));
    }
  }
};

export { moveFigure, greenBtn, diagonal, moveLeft, moveRight };
