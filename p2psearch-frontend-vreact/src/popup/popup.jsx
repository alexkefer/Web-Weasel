import React, { useState, useEffect } from "react";
import "./popup.css";

function Popup() {
  const [deviceInfo, setDeviceInfo] = useState({
    deviceName: "",
    ipAddress: "",
    nearestNode: "",
  });

  useEffect(() => {
    clearLocalStorage();
    fetchAndDisplayHostname();
    fetchAndDisplayNodeIPAddress();
    fetchAndDisplayPeersIPAddress();
  }, []);

  const clearLocalStorage = () => {
    localStorage.clear();
  }

  const fetchAndDisplayHostname = () => {
    fetch('http://localhost:8080/hostname')
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error('Failed to fetch hostname data');
        }
      })
      .then(hostname => {
        setDeviceInfo(prevState => ({
          ...prevState,
          deviceName: hostname
        }));
        localStorage.setItem('hostname', hostname);
      })
      .catch(error => {
        console.error('Error:', error.message);
      });
  };

  const fetchAndDisplayNodeIPAddress = () => {
    fetch('http://localhost:8080/peers')
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error('Failed to fetch node IP address data');
        }
      })
      .then(nodeIPAddress => {
        const lines = nodeIPAddress.split('\n');
        const firstLine = lines[0].trim();
        setDeviceInfo(prevState => ({
          ...prevState,
          ipAddress: firstLine
        }));
        localStorage.setItem('nodeIPAddress', firstLine);
      })
      .catch(error => {
        console.error('Error:', error.message);
      });
  };

  const fetchAndDisplayPeersIPAddress = () => {
    fetch('http://localhost:8080/peers')
      .then(response => {
        if (response.ok) {
          return response.text();
        } else {
          throw new Error('Failed to fetch peer IP address data');
        }
      })
      .then(nodeIPAddress => {
        const lines = nodeIPAddress.split('\n');
        const otherLines = lines.slice(1).map(line => line.trim());
        setDeviceInfo(prevState => ({
          ...prevState,
          nearestNode: otherLines.join(', ')
        }));
        localStorage.setItem('peerIPAddress', otherLines.join(', '));
      })
      .catch(error => {
        console.error('Error:', error.message);
      });
  };

  const toggleIcon = () => {
    const iconImg = document.querySelector(".icon-img");
    iconImg.src = iconImg.src.includes("on_power_icon.png")
      ? "../../images/off_power_icon.png"
      : "../../images/on_power_icon.png";

    toggleDeviceInfoVisibility(
      iconImg.src.includes("on_power_icon.png")
    );
  };

  const toggleDeviceInfoVisibility = (isIconOn) => {
    if (!isIconOn) {
      // If icon is turned off, clear all device information
      setDeviceInfo({
        deviceName: "",
        ipAddress: "",
        nearestNode: ""
      });
    } else {
      // If icon is turned on, display the previously fetched device information
      const storedHostname = localStorage.getItem('hostname');
      const storedNodeIPAddress = localStorage.getItem('nodeIPAddress');
      const storedPeerIPAddress = localStorage.getItem('peerIPAddress');
      
      setDeviceInfo({
        deviceName: storedHostname || "",
        ipAddress: storedNodeIPAddress || "",
        nearestNode: storedPeerIPAddress || ""
      });
    }
  };
  
  

  return (
    <div>
      <label className="user-info">
        <span className="device-info-text">
          Public Device Name: {deviceInfo.deviceName}
        </span>
        <span className="ip-adr-text">
          Host IP Address: {deviceInfo.ipAddress}
        </span>
        <span className="neighbor-ip-text">
          Peers IP Address: {deviceInfo.nearestNode}
        </span>
      </label>

      <label className="icon">
        <div className="icon-container">
          <button id="iconButton" className="iconButton" onClick={toggleIcon}>
            <img
              className="icon-img"
              src="../../images/on_power_icon.png"
              alt="Icon"
            />
          </button>
        </div>
      </label>

      <button
        onClick={() => window.open("webapp.html", "_blank")}
        type="button"
        className="custom-button"
      >
        Go to Web App
      </button>

      <label className="switch">
        <span className="label-text">Private Connection Request</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton1" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Hide Public Device Name</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton2" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Automatic Resource Sharing</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton3" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Automatic Resource Archiving</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton4" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>
    </div>
  );
}

export default Popup;
