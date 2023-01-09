CREATE TABLE public.players (
    id integer NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    age integer NOT NULL,
    country character varying(255) NOT NULL,
    club character varying(255) NOT NULL,
    "position" character varying(255) NOT NULL,
    goals integer NOT NULL,
    assists integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.players OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 41554)
-- Name: players_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.players_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.players_id_seq OWNER TO postgres;

--
-- TOC entry 3319 (class 0 OID 0)
-- Dependencies: 210
-- Name: players_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.players_id_seq OWNED BY public.players.id;


--
-- TOC entry 3171 (class 2604 OID 41558)
-- Name: players id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players ALTER COLUMN id SET DEFAULT nextval('public.players_id_seq'::regclass);


--
-- TOC entry 3173 (class 2606 OID 41562)
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);

CREATE SEQUENCE players_sequence
  start 1
  increment 1;
