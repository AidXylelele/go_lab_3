export class UrlService {
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