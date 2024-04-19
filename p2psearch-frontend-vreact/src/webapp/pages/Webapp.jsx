//import { useState } from 'react'
import Layout from "../Layout.jsx";

const Webapp = () => {
  return (
    <Layout>
      <div className="flex flex-col gap-y-5 ml-2">
        <h2 className={"text-2xl"}>Welcome to Peer-to-Peer Web Cache</h2>
        <p>
          This web application allows you to create a decentralized web cache
          network with your peers.
        </p>
        <div className={"flex flex-col gap-2 justify-between"}>
          <h3 className={"text-xl"}>Features:</h3>
          <ul className={"list-inside list-disc"}>
            <li>Decentralized caching of web content</li>
            <li>
              Reduced bandwidth usage by sharing cached content with peers
            </li>
            <li>Improved access speed to frequently visited websites</li>
            <li>User-friendly interface for managing cache settings</li>
          </ul>
        </div>
        <div className={"flex flex-col gap-2 justify-between"}>
          <h3 className={"text-xl"}>How to Use:</h3>
          <ol className={"list-inside list-decimal"}>
            <li>Download and install the application on your device.</li>
            <li>Create or join a peer network.</li>
            <li>Start caching and sharing web content with your peers.</li>
          </ol>
        </div>
        <p>
          For more detailed instructions, refer to the user manual or help
          section.
        </p>
        <p>Get started now and enjoy faster and more efficient web browsing!</p>
      </div>
    </Layout>
  );
};

export default Webapp;
