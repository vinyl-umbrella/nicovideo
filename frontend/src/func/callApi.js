async function search(q, page) {
  const res = await fetch(`${process.env.PREACT_APP_API_ORIGIN}/v1/search?q=${q}&page=${page}`);
  if (!res.ok) {
    console.error(res);
    throw new Error('Network response was not ok');
  }
  const json = await res.json();
  return json;
}

export default { search };
