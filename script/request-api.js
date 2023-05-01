const sendHTTPRequest = async (url) => {
  try {
    const response = await fetch(url);
    if (response.ok) return response.text();
    throw new Error(`Request failed with status ${response.status}`);
  } catch (error) {
    console.error(error);
  }
};

export { sendHTTPRequest };
