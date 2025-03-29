
type HeaderProps = {
	title: string;
}

export default function Header({ title }: HeaderProps) {
	return (
		<header className="flex justify-between items-center p-4 bg-[var(--color-tertiary)] h-[72px]">
			<h1>{title}</h1>
		</header>
	);
}