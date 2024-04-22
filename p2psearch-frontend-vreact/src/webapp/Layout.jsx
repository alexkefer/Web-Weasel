import Navigation from "./components/Navigation";
import Footer from "./components/Footer";

//eslint-disable-next-line
const Layout = ({ children }) => {
  return (
    <div
      className={
        "min-h-screen max-w-screen bg-gradient-to-br from-blue-300 to-purple-400"
      }
    >
      <div className={"flex"}>
        <div className={"h-full"}>
          <Navigation />
        </div>
        <div className={"flex flex-col mt-4 w-full"}>
          <main className={"flex-grow mx-full mx-4"}>{children}</main>
          <div className="my-auto bg-black bg-opacity-30">
            <Footer />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Layout;
