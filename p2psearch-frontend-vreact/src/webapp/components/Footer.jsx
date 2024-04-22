import { FaGithub } from "react-icons/fa";

const credits = {
  names: ["Alex Kefer", "Keagan Edwards", "Khang Tran", "Ryan Martin"],
  year: 2024,
  creator: "Made with ❤️ by",
};

const Footer = () => {
  return (
    <footer className="flex border-t-2 gap-5 justify-between self-stretch px-14 max-md:flex-wrap max-md:px-5 z-50 relative">
      <div className={"flex-auto my-auto font-bold text-violet-100"}>
        <div
          className={
            "flex flex-col gap-1 mt-1 float-left tracking-wide py-4 max:md:flex-wrap max-md:max-w-full"
          }
        >
          <h4 className={"text-md mb-2"}>{credits.creator}</h4>
          <ul className={"flex-col"}>
            {credits.names.map((name, index) => {
              return <li key={index}>{name}</li>;
            })}
          </ul>
        </div>
        <div className="flex flex-col gap-10 float-right text-xl font-bold tracking-wide uppercase text-violet-100 py-4 max-md:flex-wrap max-md:max-w-full">
          <h1 className={"text-lg justify-end"}>
            © {credits.year} - Western Washington University
          </h1>
          <div className="flex flex-row gap-5">
            <a
              href="https://www.github.com/alexkefer/p2pWebCaching"
              target="_blank"
              rel="noreferrer"
              className="text-violet-100 hover:scale-110 transform transition duration-300 ease-in-out"
            >
              <FaGithub className={"text-3xl"} />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
