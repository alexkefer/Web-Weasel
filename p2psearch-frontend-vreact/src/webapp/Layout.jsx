import Navigation from "./components/Navigation";
import Footer from "./components/Footer";

//eslint-disable-next-line
const Layout = ({ children }) => {
  return (
    <div
      className={
        "min-h-screen overflow-auto max-w-screen bg-gradient-to-br from-blue-300 to-purple-400"
      }
    >
      <div className={"flex m-auto"}>
        <div className={"h-full bg-black bg-opacity-30"}>
          <Navigation />
        </div>
        <main className={"container mx-auto mt-5"}>
          <div className={"flex-grow flex flex-col"}>{children}</div>
          <div className="w-full mt-[20em] bg-black bg-opacity-30">
            <Footer />
          </div>
        </main>
      </div>
    </div>
  );
};

export default Layout;
