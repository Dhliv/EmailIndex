import axios from "axios";

const SERVER_ADDRESS = "http://localhost:4052"

/**
 * Gets an array of emails
 * @param {int} page 
 * @returns []JSON
 */
async function getEmails(page) {
  let resp = await axios.get(`${SERVER_ADDRESS}/Emails/${page}`, {
    headers: {
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*'
    }
  });

  let jsonResp = resp.data["hits"]
  console.log(jsonResp)
  return jsonResp
}

export {getEmails}