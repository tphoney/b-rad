import { useRef, useState } from "react";
import styles from "./login.module.css";
import { authenticateUser } from "../api/user.js";
import { useSession } from "../hooks/session.js";

import Link from "../shared/link.js";

// Renders the Login page.
export default function Login({ params }) {
	const { signin, fetcher } = useSession();

	const [error, setError] = useState(null);
	const usernameElem = useRef(null);
	const passwordElem = useRef(null);

	const handleLogin = () => {
		let formData = new FormData();
		formData.append("username", usernameElem.current.value);
		formData.append("password", passwordElem.current.value);
		authenticateUser(formData, fetcher)
			.then((session) => {
				usernameElem.current.value = "";
				passwordElem.current.value = "";
				signin(session);
			})
			.catch((error) => {
				setError(error);
			});
	};

	const alert =
		error && error.message ? (
			<div class="alert">{error.message}</div>
		) : undefined;

	return (
		<>
			<section className={styles.root}>
				<h2>Login</h2>
				{alert}
				<input
					ref={usernameElem}
					type="text"
					name="username"
					placeholder="Email"
				/>
				<input
					ref={passwordElem}
					type="password"
					name="password"
					placeholder="Password"
				/>
				<button onClick={handleLogin}>Login</button>
				<div class="actions">
					<Link href="/register">Create a new account</Link>
				</div>
			</section>
		</>
	);
}
