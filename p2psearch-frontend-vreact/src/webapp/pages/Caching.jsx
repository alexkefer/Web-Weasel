import React, { useState, useEffect } from "react";
import Layout from "../Layout.jsx";

const Caching = () => {
  // State for storing URL list
  const [urlList, setUrlList] = useState([]);
  // State for input value
  const [urlInput, setUrlInput] = useState("");
  // State for output message
  const [outputMessage, setOutputMessage] = useState("");

  // Retrieve stored URL list from localStorage on component mount
  useEffect(() => {
    const storedUrlList = JSON.parse(localStorage.getItem("urlList"));
    if (storedUrlList) {
      setUrlList(storedUrlList);
    }
  }, []);

  // Function to update and store the URL list in localStorage
  const updateAndStoreUrlList = () => {
    localStorage.setItem("urlList", JSON.stringify(urlList));
  };

  // Function to clear a specific URL from the list
  const clearUrl = (url) => {
    const updatedUrlList = urlList.filter((u) => u !== url);
    setUrlList(updatedUrlList);
    updateAndStoreUrlList();
  };

  // Function to handle fetch button click
  const handleFetchButtonClick = () => {
    const urlRegex = /^(https?|http):\/\/[^\s$.?#].[^\s]*$/i;

    if (!urlInput || !urlRegex.test(urlInput)) {
      setOutputMessage("Please enter a valid URL.");
      return;
    }

    if (urlList.includes(urlInput)) {
      setOutputMessage("URL is already in the list.");
      return;
    }

    const updatedUrlList = [...urlList, urlInput];
    setUrlList(updatedUrlList);
    setUrlInput(""); // Clear input field

    updateAndStoreUrlList();

    const url = "http://localhost:8080/cache?path=" + encodeURIComponent(urlInput);

    fetch(url, { mode: "no-cors" })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.text();
      })
      .then((data) => {
        console.log("Cached URL:", data);
        setOutputMessage("Cached URL: " + data);
      })
      .catch((error) => {
        console.error("There was a problem with your fetch operation:", error);
      });
  };

  return (
    <Layout>
      <div id="downloadContainer">
        <h1>Web Caching</h1>
        <label htmlFor="urlInput">Enter URL:</label>
        <input
          type="text"
          id="urlInput"
          value={urlInput}
          onChange={(e) => setUrlInput(e.target.value)}
        />
        <button id="fetchButton" onClick={handleFetchButtonClick}>
          Fetch
        </button>
        <p id="output">{outputMessage}</p>
        <h2>Saved URL</h2>
        <ul id="urlList">
          {urlList.map((url) => (
            <li key={url}>
              <a href={"http://localhost:8080/retrieve?path=" + url} target="_blank" rel="noopener noreferrer">
                {url}
              </a>
              <button onClick={() => clearUrl(url)}>Clear</button>
            </li>
          ))}
        </ul>
      </div>
    </Layout>
  );
};

export default Caching;
