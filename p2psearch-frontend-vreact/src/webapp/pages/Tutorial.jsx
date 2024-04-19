import Layout from "../Layout.jsx";

const Tutorial = () => {
  return (
    <Layout>
      <div className="content-section">
        <h2>Tutorial</h2>
        <p>
          This web application allows you to create a decentralized web cache
          network with your peers.
        </p>
        <h3>How to Install:</h3>
        <ol>
          <li>Clone the github repo to your local device.</li>
          <li>Open your chromium based browser and go to extensions.</li>
          <li>Turn on developer mode.</li>
          <li>Click load unpacked.</li>
          <li>Navigate to ./p2pWebCaching/p2psearch-frontend</li>
          <li>Click open.</li>
          <p>
            You now should be able to open the extension and the website in your
            browser
          </p>
        </ol>
        <h3>How to Connect:</h3>
        <ol>
          <li>Open up your console.</li>
          <li>Navigate to ./p2pWebCaching/p2psearch-backend</li>
          <li>
            Enter the command <code>go run .</code> if you are <u>hosting</u>,
            or <code>go run . ip &lt;addr&gt;:&lt;port&gt;</code> if you are{" "}
            <u>connecting</u> to an instance.
          </li>

          <p>
            To find the port and ip, ask for them from the host you want to
            connect to. They should be in lines below the <code>go run .</code>{" "}
            command on the host's console. These will be marked with the{" "}
            <code className={"text-blue-600"}>[INFO]</code> tag.
          </p>
        </ol>
        <h3>How to Use:</h3>
        <ol>
          <li>
            Run the go server using <code>go run .</code> in the console when
            navigated to ./p2pWebCaching/p2psearch-backend
          </li>
          <li>
            Choose a website url you wish to access through the p2p extension.
          </li>
          <li>Navigate to the Caching page.</li>
          <li>Paste the URL into the fetch text box.</li>
          <li>Click fetch.</li>
          <p>
            The URL should now be a clickable link for all peers and viewable.
            Once clicked, the URL on that page when visiting through the
            extension should be in similar format to "http://localhost:8080"
          </p>
        </ol>
      </div>
    </Layout>
  );
};

export default Tutorial;
