import React from "react";

export class API {
  constructor(basePath, getState, setState) {
    this.basePath = basePath;
    this.getState = getState;
    this.setState = setState;
  }

  checkFetchAndSet(key, endpoint) {
    let value = this.getState(key);
    if (!value && !this.fetching) {
      this.fetching = true;
      fetch(this.basePath + endpoint)
        .then(response => response.json())
        .then(json => {
          this.setState(key, json);
          value = json;
          this.fetching = false;
        })
        .catch(error => {
          console.error(error);
          this.fetching = false;
        });
    }

    return value;
  }

  getNumBlogs() {
    let numBlogs = this.checkFetchAndSet("numBlogs", "blogs/count");
    return numBlogs ? numBlogs.count : 0;
  }

  getLatestPosts(lang = "en") {
    return this.checkFetchAndSet("posts-latest-" + lang, "latest/" + lang);
  }
}

export const APIContext = React.createContext(new API());
