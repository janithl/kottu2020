import React, { useState } from "react";

function SearchBar() {
  const [query, setQuery] = useState("");

  return (
    <form className="navbar-form navbar-left" method="GET" action="/search/">
      <div className="form-group">
        <input
          tabIndex="1"
          id="searchbar"
          name="q"
          type="text"
          className="form-control"
          placeholder="Search Kottu..."
          value={query}
          onChange={e => setQuery(e.target.value)}
        />
      </div>
      <button
        tabIndex="2"
        id="searchbtn"
        type="submit"
        className="btn btn-default"
      >
        Search
      </button>
    </form>
  );
}

export default SearchBar;
