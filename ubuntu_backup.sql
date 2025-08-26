--
-- PostgreSQL database dump
--

-- Dumped from database version 16.9 (Ubuntu 16.9-0ubuntu0.24.04.1)
-- Dumped by pg_dump version 16.9 (Ubuntu 16.9-0ubuntu0.24.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: blogs; Type: TABLE; Schema: public; Owner: ubuntu
--

CREATE TABLE public.blogs (
    id integer NOT NULL,
    blog_subject character varying(255) NOT NULL,
    blog_owner_name character varying(255) NOT NULL,
    blog_owner_email character varying(255) NOT NULL,
    blog_body text NOT NULL,
    blog_category character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.blogs OWNER TO ubuntu;

--
-- Name: blogs_id_seq; Type: SEQUENCE; Schema: public; Owner: ubuntu
--

CREATE SEQUENCE public.blogs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blogs_id_seq OWNER TO ubuntu;

--
-- Name: blogs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ubuntu
--

ALTER SEQUENCE public.blogs_id_seq OWNED BY public.blogs.id;


--
-- Name: comments; Type: TABLE; Schema: public; Owner: ubuntu
--

CREATE TABLE public.comments (
    id integer NOT NULL,
    comment_blog_id integer NOT NULL,
    comment_name character varying(255) NOT NULL,
    comment_email character varying(255) NOT NULL,
    comment_body text NOT NULL,
    comment_approved boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.comments OWNER TO ubuntu;

--
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: ubuntu
--

CREATE SEQUENCE public.comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.comments_id_seq OWNER TO ubuntu;

--
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ubuntu
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- Name: places; Type: TABLE; Schema: public; Owner: ubuntu
--

CREATE TABLE public.places (
    id integer NOT NULL,
    place_name character varying(255) NOT NULL,
    place_info text,
    place_lat numeric(10,8),
    place_lng numeric(11,8),
    place_icon_type integer DEFAULT 1,
    place_address text,
    place_phone character varying(50),
    place_email character varying(255),
    place_website character varying(255),
    place_arrive timestamp without time zone,
    place_depart timestamp without time zone,
    place_hide_info boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.places OWNER TO ubuntu;

--
-- Name: places_id_seq; Type: SEQUENCE; Schema: public; Owner: ubuntu
--

CREATE SEQUENCE public.places_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.places_id_seq OWNER TO ubuntu;

--
-- Name: places_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ubuntu
--

ALTER SEQUENCE public.places_id_seq OWNED BY public.places.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: ubuntu
--

CREATE TABLE public.users (
    id integer NOT NULL,
    user_name character varying(255) NOT NULL,
    user_password character varying(255) NOT NULL,
    user_email character varying(255) NOT NULL,
    user_role character varying(50) DEFAULT 'User'::character varying,
    user_approved boolean DEFAULT false,
    user_verify_code text,
    user_verify_expires timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO ubuntu;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: ubuntu
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO ubuntu;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ubuntu
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: blogs id; Type: DEFAULT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.blogs ALTER COLUMN id SET DEFAULT nextval('public.blogs_id_seq'::regclass);


--
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- Name: places id; Type: DEFAULT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.places ALTER COLUMN id SET DEFAULT nextval('public.places_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: blogs; Type: TABLE DATA; Schema: public; Owner: ubuntu
--

COPY public.blogs (id, blog_subject, blog_owner_name, blog_owner_email, blog_body, blog_category, created_at, updated_at) FROM stdin;
1	Troubleshoot an Aquahot	ecsproull	ecsproull@outlook.com	<p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">Troubleshooting an Aqua-Hot system in an RV can help resolve many common issues without a service call. Aqua-Hot systems are used for hydronic heating and hot water in many high-end RVs. Here's a list of common tips for troubleshooting: 🔧 Basic Checks (First Things to Inspect)</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);"><b style="box-sizing: inherit; font-weight: 700;">Power Supply</b>: Ensure the main power switch (12V DC) is ON. Check the house batteries — low voltage can cause system failure. Confirm circuit breakers and fuses for the Aqua-Hot system are not tripped or blown.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);"><b style="box-sizing: inherit; font-weight: 700;">Fluid Level:</b><span>&nbsp;</span>Check the antifreeze solution (propylene glycol) level in the reservoir. It should be between "COLD" and "HOT" marks when the unit is cold. Low fluid can cause no heat or burner lockout.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);"><b style="box-sizing: inherit; font-weight: 700;">Thermostats:</b><span>&nbsp;</span>Make sure thermostats inside the coach are set higher than the current room temperature. Test each zone independently if possible.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);"><b style="box-sizing: inherit; font-weight: 700;">Diesel Burner:</b><span>&nbsp;</span>Listen for the burner ignition; if you hear clicking but no flame, it might be a fuel or ignition issue. Check the diesel fuel level (some units won't operate below 1/4 tank). Inspect the fuel filter or nozzle if it hasn't been serviced recently.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">🔥<span>&nbsp;</span><b style="box-sizing: inherit; font-weight: 700;">Heating Issues:</b>&nbsp;Of there is no heat from furnace fans, check if the fans are running — they require 12V power. If the fans run but no warm air, the heat exchanger may be cold — indicating a burner or electric element issue.&nbsp;</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">If you have heat from electric element only and you're plugged into shore power and only getting warm water, your electric element might be working, but not the burner. The burner may need cleaning, service, or fuel supply checked. You must have above a quarter tank of fuel for the burner to ignite.&nbsp;</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">If you have inconsistent Interior heating, make sure all heat zones are working. Some zones may have air locks or faulty thermostats. Bleed the system if there's suspected air in the hydronic lines.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">💧<span>&nbsp;</span><b style="box-sizing: inherit; font-weight: 700;">Hot Water Issues:</b><span>&nbsp;</span>No Hot Water Ensure the Domestic Water Loop switch is ON. Check if bypass valves are incorrectly set (from winterizing). Confirm if the Aqua-Hot system is reaching temperature (160–180°F typically).</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);"><b>Inconsistent Hot Water:</b> The mixing valve might be faulty or clogged. Also, ensure enough time has passed for the burner or electric element to heat water.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">🚨 <b>Error or Lockout Conditions Red Lockout Light (Diesel Burner)</b>: Often indicates a failed ignition or fuel issue. Reset the system: Turn the burner OFF for 30 seconds and back ON. If lockout persists, inspect for: Fuel delivery problems Faulty ignition coil or clogged nozzle.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);"><b>Burner Won’t Start:</b> Check fuel level, inline filter, and clean the nozzle. Test the photo eye (may need cleaning if covered in soot). Testing the photo eye involves disassembling the Aqua-hot. This is covered in they rebuilding an Aqua-hot blog. (not created yet)</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">🧰 Maintenance Tips Change fuel filter and nozzle annually. Inspect and clean burner and photo eye yearly. Flush and refill coolant every 2–3 years. Bleed air from the heating loop after any major service.</p><p class="MuiTypography-root MuiTypography-body1 MuiTypography-paragraph css-rn0nqi-MuiTypography-root" style="box-sizing: inherit; margin: 0px 0px 16px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial; line-height: 1.5; background-color: rgb(238, 238, 238);">🛠 <b>Tools to Keep Handy</b>: A multi-meter for 12V power testing, wrenches for nozzle or burner access, a flashlight and cleaning supplies for photo eye or soot.</p>	RV Technology	2025-07-13 00:16:28.918	2025-07-13 01:09:18.153
2	How GitHub Copilot Changed the Way I Code	ecsproull	ecsproull@outlook.com	<p data-start="311" data-end="636">After decades of wrestling with code, learning new frameworks from books, and debugging until my eyes crossed, I decided to try something completely different for my latest React project: <strong data-start="499" data-end="517">GitHub Copilot</strong>. I've been studying programming technologies for many years, but this was my first time using AI as my coding partner.</p>\n<p data-start="638" data-end="697">What an eye-opening experience it's been—both good and bad.</p>\n<hr data-start="699" data-end="702">\n<h2 data-start="704" data-end="745">The Project: Tailwind to MUI Migration</h2>\n<p data-start="747" data-end="947">My goal was to take an existing React project, this website, that used Tailwind CSS and convert it to <strong data-start="834" data-end="855">MUI (Material UI)</strong>—a library I had never used before. Actually, until this project, I hadn’t even heard of it.</p>\n<p data-start="949" data-end="1173">I began by creating a new React project and installing the necessary Node modules for MUI. Once the basic structure was in place, I imported one React component at a time and asked Copilot to convert it from Tailwind to MUI.</p>\n<p data-start="1175" data-end="1392">The first few components worked surprisingly well, though there were always a few mistakes. The biggest surprise? Copilot would sometimes generate broken code—then immediately help fix the very problems it introduced.</p>\n<hr data-start="1394" data-end="1397">\n<h2 data-start="1399" data-end="1433">A Career’s Worth of Perspective</h2>\n<p data-start="1435" data-end="1785">Over the course of my software development career—about 40 years—I’ve seen huge changes in how we build and learn software. When I started at Microsoft in 2000, most of my time was spent digging through large codebases. As search engines got better, I found that some of the best insights into Windows internals came from a well-formed Google search.</p>\n<p data-start="1787" data-end="1913">But it was always a juggling act: jumping between a browser, Visual Studio, WinDbg, and whatever docs I could scrape together.</p>\n<p data-start="1915" data-end="1987">Now, with tools like GitHub Copilot, that fragmentation is disappearing.</p>\n<hr data-start="1989" data-end="1992">\n<h2 data-start="1994" data-end="2014">AI Inside the IDE</h2>\n<p data-start="2016" data-end="2265">One of the most transformative aspects of Copilot is that it lives inside your IDE. It has direct access to your code, so there's no more context-switching. Just highlight a block of code—or even place your cursor near it—and start asking questions.</p>\n<p data-start="2267" data-end="2456">It will do its best to analyze the surrounding context, generate or modify code, and even explain its reasoning. Often, it offers multiple approaches, with clear pros and cons for each one.</p>\n<p data-start="2458" data-end="2554">That kind of inline, contextual support would have seemed like science fiction not too long ago.</p>\n<hr data-start="2556" data-end="2559">\n<h2 data-start="2561" data-end="2582">Ditching the Books</h2>\n<p data-start="2584" data-end="2872">I used to rely heavily on programming books. For this project, I found a book on MUI written around 2020—but the examples wouldn’t even run. I copied snippets from the book and asked Copilot to update them to the current version. It did such a good job that I abandoned the book entirely.</p>\n<p data-start="2874" data-end="3113">If I had stuck with the book, it probably would’ve taken a month or two just to get comfortable enough to start the conversion. With Copilot, I was able to build and deploy the entire site—including this blog post—in less than three weeks.</p>\n<hr data-start="3115" data-end="3118">\n<h2 data-start="3120" data-end="3152">Is It Perfect? Not Even Close</h2>\n<p data-start="3154" data-end="3284">Was it flawless? Definitely not. But I don’t recall running into any issue that I couldn’t debug and fix. And that’s saying a lot.</p>\n<p data-start="3286" data-end="3538">There’s plenty of talk about AI replacing programmers. Will it change the industry? Absolutely. The skill set required to be a developer will evolve, and quickly. But the fundamentals of software design and problem-solving? Those aren’t going anywhere.</p>\n<hr data-start="3540" data-end="3543"><br>	Software	2025-07-13 03:52:41.84	2025-07-13 03:52:41.84
3	Ride Around Fidalgo Island	ecsproull	ecsproull@outlook.com	<p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><span style="box-sizing: inherit; font-size: 1rem; letter-spacing: 0.00938em;">Today’s workout was an easy ride for an hour, so I decided to do a couple of laps around the refineries on Fidalgo Island. While it is called an island, it is easily accessible by several roads. The Tommy Thomson Trail runs directly through our RV park and ends on the loop I wanted to ride. Two laps take almost exactly an hour for me, but today I decided to stop and take a few pictures. First up is a lot of water with Mt. Baker in the background. It would probably be fair to point out that right behind me is a huge Marathon Refinery.</span></p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><img src="fidalgo_bike/mt_baker.webp"></p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;">Almost immediately after putting my phone away, I saw a large eagle heading toward the side of the road that I was riding on and coming in for a landing. Within seconds, it snatched up a small snake and was off again. I was really bummed that I missed that as it would have been an awesome picture.</p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;">Further up the road, there is a pasture with a herd of beef cows. They always look curious when I ride by, and I joke that I make so many loops past them that we are on a first-name basis. As I stopped to take a picture, I noticed one younger calf that appeared to be curious and started walking toward me.</p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><img src="fidalgo_bike/cows.webp"></p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;">I mentioned the Tommy Thomson Trail runs through the RV park, and I use that to cross Fidalgo Bay to get to the bike loop. This trail used to be a railroad passage long ago. I am really curious if this railroad car was part of the railroad that used to pass through here. It is fenced off and sits in the middle of a field that I pass by often. I have no idea how it got there or why it is here, but it certainly speaks of a time gone by.</p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><img src="fidalgo_bike/traincar2.webp"></p><p style="box-sizing: inherit; margin-bottom: 8px; color: rgba(0, 0, 0, 0.87); font-family: Roboto, Helvetica, Arial, sans-serif; font-size: 16px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: 0.15008px; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(238, 238, 238); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;">All in all, it was just another day of training near Anacortes, WA. I finished up by making a short trek out to the casino on Rt 20 and completed two loops around the refineries.</p>	Triathlon & Training	2025-07-31 01:13:16.685	2025-07-31 01:17:12.219
4	AI Coding: Why Experience Still Matters	ecsproull	ecsproull@outlook.com	<p data-start="225" data-end="690">I’ve been writing software for over forty years, starting with BASIC on an IBM PC clone in the early 1980s. In the mid-1990s I went back to college, then spent time working with C++/MFC at Rockwell Automation before joining Microsoft in 2000. During my years there I worked with just about every technology the company had to offer. I retired in September 2017, but now, in my 70s, I still write software for fun and to keep up with the rapid changes in our industry.</p>\n<p data-start="692" data-end="978">Over those decades, I’ve seen it all—mainframes, minicomputers, desktop PCs, client/server, the web, cloud computing. Every new wave of technology promised to “change everything,” and many of them did. But nothing has arrived with the speed and disruptive force of AI-assisted coding.</p>\n<p data-start="980" data-end="1308">When I first tried AI for code generation, it felt like being handed the keys to a high-performance sports car. The raw power was intoxicating. The acceleration was breathtaking. And unlike learning a new language or framework, the ramp-up was practically zero: describe what you want, and seconds later you have working code.</p>\n<p data-start="1310" data-end="1500">But here’s the thing: I’ve been driving for decades. I know when to lean into the curves and when to tap the brakes. I don’t panic when I see a road hazard, because I’ve been there before.</p>\n<p data-start="1502" data-end="1884">The same applies to AI-assisted coding. AI can generate a working solution instantly, but “working” and “correct” are not the same thing. “Working” and “secure” are definitely not the same thing. AI doesn’t have judgment. It doesn’t know the business context, the architectural constraints, the performance limits, or the compliance requirements. It just knows what usually works.</p>\n<p data-start="1886" data-end="1937">That’s where experience makes all the difference.</p>\n<p data-start="1939" data-end="2007">I can look at AI-generated code and spot the pitfalls immediately:</p>\n<ul data-start="2008" data-end="2212"><li data-start="2008" data-end="2080">\n<p data-start="2010" data-end="2080">This function looks fine, but it will blow up under real-world load.</p>\n</li><li data-start="2081" data-end="2146">\n<p data-start="2083" data-end="2146">That loop works, but it won’t scale—there’s a better pattern.</p>\n</li><li data-start="2147" data-end="2212">\n<p data-start="2149" data-end="2212">This SQL query runs, but it’s wide open to injection attacks.</p>\n</li></ul>\n<p data-start="2214" data-end="2699">Now imagine handing this same tool to someone who just graduated. A new engineer, fresh from a university program that may or may not reflect industry realities, suddenly has the ability to churn out “complete” applications at speed. Will they understand what’s happening under the hood? Will they know when the AI’s solution is subtly wrong? Or is this like handing the keys to a 700-horsepower sports car to someone who just got their license and saying, “Have fun on the freeway”?</p>\n<p data-start="2701" data-end="2885">This is not a knock on new engineers—we all start somewhere. The difference is that my early mistakes happened at 40 mph in an empty parking lot, not at 200 mph in rush-hour traffic.</p>\n<p data-start="2887" data-end="2910">So what’s the answer?</p>\n<p data-start="2912" data-end="3321">In the AI era, mentorship will matter more than ever. New developers need guidance not just in syntax and frameworks, but in architecture, testing, debugging, and security. Schools should teach how to <em data-start="3113" data-end="3120">think</em> about code, not just how to type it. And as an industry, we need to double down on pairing, reviews, and explanation—helping the next generation understand not just <em data-start="3286" data-end="3292">what</em> works, but <em data-start="3304" data-end="3309">why</em> it works.</p>\n<p data-start="3323" data-end="3560">AI is here, and it’s not going away. Used wisely, it’s a game changer. Used blindly, it’s a dangerous tool. The solution isn’t to keep the sports car locked in the garage—it’s to make sure the driver’s education matches the horsepower.</p>\n<p data-start="3562" data-end="3725">And as someone who’s been driving for a very long time, I can tell you this: the ride has never been more exciting—or more in need of a steady hand on the wheel.</p>	Software	2025-08-26 00:18:50.639935	2025-08-26 00:18:50.639935
\.


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: ubuntu
--

COPY public.comments (id, comment_blog_id, comment_name, comment_email, comment_body, comment_approved, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: places; Type: TABLE DATA; Schema: public; Owner: ubuntu
--

COPY public.places (id, place_name, place_info, place_lat, place_lng, place_icon_type, place_address, place_phone, place_email, place_website, place_arrive, place_depart, place_hide_info, created_at, updated_at) FROM stdin;
1	Home		33.66419260	-112.37534900	2	19819 N White Rock Dr	4253513207	ecs3@po.cwru.edu	https://edandlinda.com	2025-05-01 00:00:00	2025-05-13 00:00:00	t	2025-08-26 00:18:50.721989	2025-08-26 00:18:50.721989
2	J & J RV Park	One night stop	37.04241410	-112.51722230	1	584 East 300 South Kanab, UT 84741	 435-899-1956	jandjrvpark@gmail.com	https://jandjrvpark.com/	2025-05-13 00:00:00	2025-05-14 00:00:00	f	2025-08-26 00:18:50.760329	2025-08-26 00:18:50.760329
3	Love's RV Stop		41.50886690	-112.05406440	1	20 South 1550 West Street Brigham City, UT 84302	435-921-9983	store686@loves.com	www.lovesrvstops.com	2025-05-14 00:00:00	2025-05-15 00:00:00	f	2025-08-26 00:18:50.761814	2025-08-26 00:18:50.761814
4	Deer Lodge A-OK RV park		46.39882110	-112.74075250	1	330 Park St, Deer Lodge, MT 	406-846-1629	Deerlodgeaokcampground@gmail.com	https://deerlodgeaokcampground.com/	2025-05-15 00:00:00	2025-05-16 00:00:00	f	2025-08-26 00:18:50.764006	2025-08-26 00:18:50.764006
5	Blackwell Island RV Resort	Will be for IMCda 70.3 and then leave for Salem OR for another 70.3	47.68256000	-116.80230000	1	800 S Marina Dr, Coeur d'Alene ID	208-930-1865	stay@idahorvpark.com	idahorvpark.com	2025-05-16 00:00:00	2025-07-17 00:00:00	f	2025-08-26 00:18:50.843091	2025-08-26 00:18:50.843091
6	Fidalgo Bay Rv Resort		48.48208930	-122.59354470	1	4701 Fidalgo Bay Rd Anacortes WA 98221	360-293-5353	Reserve@FidalgoBay.com	https://fidalgobay.com/	2025-07-17 00:00:00	2025-09-01 00:00:00	f	2025-08-26 00:18:50.880621	2025-08-26 00:18:50.880621
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: ubuntu
--

COPY public.users (id, user_name, user_password, user_email, user_role, user_approved, user_verify_code, user_verify_expires, created_at, updated_at) FROM stdin;
1	ecsproull	$2b$10$GlodPQ7Zyfu66ICrSnrVXu8UGT9eP44Wq7cTowOc58Au/pAop89.i	ecsproull@outlook.com	Admin	t		0001-01-01 00:00:00	2025-08-26 00:18:50.078358	2025-08-26 00:18:50.078358
2	ecsproull3	$2b$10$.5TdOQw8Hhl8OM9JIZrI/.EL6E5wuRZN1Io7reSvu.IyjjYWSFmuy	ecs3@po.cwru.edu	Commentor	t		0001-01-01 00:00:00	2025-08-26 00:18:50.121213	2025-08-26 00:18:50.121213
\.


--
-- Name: blogs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ubuntu
--

SELECT pg_catalog.setval('public.blogs_id_seq', 4, true);


--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ubuntu
--

SELECT pg_catalog.setval('public.comments_id_seq', 1, false);


--
-- Name: places_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ubuntu
--

SELECT pg_catalog.setval('public.places_id_seq', 6, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ubuntu
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: blogs blogs_pkey; Type: CONSTRAINT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.blogs
    ADD CONSTRAINT blogs_pkey PRIMARY KEY (id);


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: places places_pkey; Type: CONSTRAINT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.places
    ADD CONSTRAINT places_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ubuntu
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

