export default function timeDiff(timestamp: string): string {
  const past = new Date(timestamp);
  const now = new Date();

  const diffMs = now.getTime() - past.getTime();

  const minutes = diffMs / (1000 * 60);
  const hours = diffMs / (1000 * 60 * 60);
  const days = diffMs / (1000 * 60 * 60 * 24);

  if (minutes < 60) {
    return `${Math.floor(minutes)} minute(s) ago`;
  } else if (hours < 24) {
    return `${Math.floor(hours)} hour(s) ago`;
  } else {
    return `${Math.floor(days)} day(s) ago`;
  }
}
