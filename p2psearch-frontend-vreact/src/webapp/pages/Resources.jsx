import Layout from "../Layout.jsx";

const Resources = () => {
  return (
    <Layout>
      <section className="content-section">
        <h2>GitHub Repositories</h2>
        <p>
          Visit our GitHub repositories to access additional resources,
          documentation, and contribute to the project:
        </p>
        <ul>
          <li>
            <a
              href="https://github.com/alexkefer/p2pWebCaching"
              target="_blank"
            >
              Peer-to-Peer Networking Component
            </a>
          </li>
          <li>
            <a
              href="https://pkg.go.dev/github.com/alexkefer/webDownloader"
              target="_blank"
            >
              Web Caching Component
            </a>
          </li>
        </ul>
      </section>
    </Layout>
  );
};

export default Resources;
