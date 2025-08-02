// The example file for testing

const html5QrCode = new Html5Qrcode("reader");

let canScan = true;

function onScanSuccess(decodedText, decodedResult) {
  if (!canScan) return;

  canScan = false;

  console.log(`Сканировано: ${decodedText}`);

  const parts = decodedText.split(";");
  const data = {};
  parts.forEach((part) => {
    const [key, value] = part.split(":");
    if (key && value) {
      data[key.trim()] = value.trim();
    }
  });

  const params = new URLSearchParams(window.location.search);

  const phone = data.phone || data.mobile_phone;
  const token = data.mobile || data.token;
  const seller_chat_id = Number(params.get("chat_id"));
  const seller_id = Number(params.get("user_id"));

  if (!phone || !token) {
    console.log("error: qr need to contains phone and token");
  } else {
    sendToServer(phone, token, seller_chat_id, seller_id);
  }

  setTimeout(() => {
    canScan = true;
  }, 1000);
}

async function sendToServer(phone, token, seller_chat_id, seller_id) {
  const url = "https://thin-papers-mix.loca.lt/auth";
  const payload = { phone, token, seller_chat_id, seller_id };

  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.text();
    console.log("Success send:", result);
  } catch (error) {
    console.error("error while sending:", error);
  }
}

function onScanFailure(error) {}

window.addEventListener("load", () => {
  const config = {
    fps: 10,
    qrbox: { width: 250, height: 250 },
    aspectRatio: 1.0,
    rememberLastUsedCamera: true,
  };

  html5QrCode
    .start(
      { facingMode: { exact: "environment" } },
      config,
      onScanSuccess,
      onScanFailure,
    )
    .catch((err) => {
      console.warn(
        "Тыловая камера недоступна, переключение на фронтальную:",
        err,
      );

      html5QrCode
        .start({ facingMode: "user" }, config, onScanSuccess, onScanFailure)
        .catch((err) => {
          console.error("Ошибка при запуске сканирования:", err);
        });
    });
});

