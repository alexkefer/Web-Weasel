import { useState, useEffect } from "react";
import Layout from "../Layout.jsx";

const Caching = () => {
  // State for storing URL list
  const [urlList, setUrlList] = useState([]);
  // State for input value
  const [urlInput, setUrlInput] = useState("");
  // State for output message
  const [outputMessage, setOutputMessage] = useState("");

  // Function to retrieve stored URL list from the server on page load
  useEffect(() => {
    fetch("http://localhost:8080/sites")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.text(); // Get the response as plain text
      })
      .then((text) => {
        const storedUrlList = text
          .split("\n")
          .filter((url) => url.trim() !== "");
        if (storedUrlList.length > 0) {
          setUrlList(storedUrlList);
        }
      })
      .catch((error) => {
        console.error(
          "There was a problem with fetching the URL list:",
          error
        );
      });
  }, []);

  // Function to clear a specific URL from the list and update the display
  const clearUrl = (url) => {
    const trimmedUrl = url.replace(/^https?:\/\//, "");
    fetch("http://localhost:8080/removeSite?url=" + trimmedUrl, {
      method: "GET", // or 'DELETE' if supported by the server
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        // Remove the URL from the local list and update the display
        const updatedUrlList = urlList.filter((u) => u !== url);
        setUrlList(updatedUrlList);
        // Display output message
        setOutputMessage("URL successfully removed: " + url);
      })
      .catch((error) => {
        console.error("There was a problem with removing the URL:", error);
      });
  };

  // Function to handle fetch button click
  const handleFetchButtonClick = () => {
    let urlInput = document.getElementById("urlInput").value.trim();

    // Regular expression to match URL pattern
    const urlRegex = /^(https?|http):\/\/[^\s$.?#].[^\s]*$/i;

    // Automatically prepend "https://" if missing
    if (!urlInput.startsWith("http://") && !urlInput.startsWith("https://")) {
      urlInput = "https://" + urlInput;
    }

    // Validate the URL
    if (!urlRegex.test(urlInput)) {
      setOutputMessage("Please enter a valid URL.");
      return;
    }

    // Check if the URL is already in the list
    if (urlList.includes(urlInput)) {
      setOutputMessage("URL is already in the list.");
      return;
    }

    // Push the valid URL to the urlList array
    setUrlList([...urlList, urlInput]);
    setUrlInput(""); // Clear input field

    const url =
      "http://localhost:8080/cache?path=" + encodeURIComponent(urlInput);

    fetch(url, { mode: "no-cors" })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.text();
      })
      .then((data) => {
        console.log("Cached URL:", data);
        setOutputMessage("Cached URL: " + data); // Set the message
      })
      .catch((error) => {
        console.error(
          "There was a problem with your fetch operation:",
          error
        );
      });
  };

  return (
    <Layout>
      <div
        id="downloadContainer"
        className="flex flex-col gap-y-5 max-w-3xl max-lg:max-w-lg mx-auto"
      >
        <h1 className={"text-3xl"}>Web Caching</h1>
        <div className={"input-container"}>
          <label htmlFor="urlInput">Enter web page URL:</label>
          <input
            type="text"
            id="urlInput"
            value={urlInput}
            onChange={(e) => setUrlInput(e.target.value)}
          />
          <button id="fetchButton" onClick={handleFetchButtonClick}>
            Download
          </button>
        </div>
        <p id="output" className={"response-message"}>
          {outputMessage}
        </p>
        <div className={"saved-pages"}>
          <h2 className={"text-2xl"}>Saved Webpages</h2>
          <p>
            Click on the URL to view the cached webpage. Click "Clear" to remove
            the URL from the list.
          </p>
          <ul id="urlList">
            {urlList.map((url) => (
              <li key={url}>
                <a
                  href={"http://localhost:8080/retrieve?path=" + url}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  {url.replace(/^https?:\/\//, "")}
                </a>
                <button onClick={() => clearUrl(url)}>Clear</button>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </Layout>
  );
};

export default Caching;
