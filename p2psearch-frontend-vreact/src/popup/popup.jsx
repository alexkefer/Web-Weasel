function Popup() {
    return (
      <div>
        <label className="user-info">
          <span className="device-info-text">Public Device Name</span>
          <span className="ip-adr-text">Node IP Addresses</span>
          <span className="neighbor-ip-text">Nearest Connection Node</span>
        </label>
  
        <label className="icon">
          <div className="icon-container">
            <button id="iconButton" className="iconButton">
              <img className="icon-img" src="../images/on_power_icon.png" alt="Icon" />
            </button>
          </div>
        </label>
  
        <button onClick={() => window.open('webapp.html', '_blank')} type="button" className="custom-button">
          <a href="webapp.html" target="_blank">Go to Web App</a>
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
