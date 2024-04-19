import Layout from "../Layout.jsx";

const Settings = () => {
  return (
    <Layout>
      <section className="settings-section">
        <h2>Settings</h2>

        <div className="setting-container">
          <span className="label-text">Request Private Connections</span>
          <div className="slider-container">
            <input
              type="checkbox"
              id="toggleButton1"
              className="switch-input"
            />
            <span className="slider"></span>
          </div>
        </div>
        <p className="setting-description">
          When enabled, only accept connection requests from approved users.
        </p>

        <div className="setting-container">
          <span className="label-text">Hide Public Device Name</span>
          <div className="slider-container">
            <input
              type="checkbox"
              id="toggleButton2"
              className="switch-input"
            />
            <span className="slider"></span>
          </div>
        </div>
        <p className="setting-description">
          Hide your device's name from being visible to others on the network.
        </p>

        <div className="setting-container">
          <span className="label-text">Enable Automatic Resource Sharing</span>
          <div className="slider-container">
            <input
              type="checkbox"
              id="toggleButton3"
              className="switch-input"
            />
            <span className="slider"></span>
          </div>
        </div>
        <p className="setting-description">
          Automatically share resources with trusted devices on the network.
        </p>

        <div className="setting-container">
          <span className="label-text">
            Enable Automatic Resource Archiving
          </span>
          <div className="slider-container">
            <input
              type="checkbox"
              id="toggleButton4"
              className="switch-input"
            />
            <span className="slider"></span>
          </div>
        </div>
        <p className="setting-description">
          Automatically archive resources after sharing them with other devices.
        </p>
      </section>
    </Layout>
  );
};

export default Settings;
