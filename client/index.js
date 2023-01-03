function test() {
  alert("Hi ok");
}
async function subscribe() {
  navigator.serviceWorker.ready
    .then(function (registration) {
      const vapidPublicKey =
        "BIZoty0iCIkAuR4rZtLEaixmbIL3-Xwg3NO8BkX6-vrcIX0UHdZ8Ux9JojPhA8AutoUNANGPHxayJiv4IKcM8po";

      return registration.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(vapidPublicKey),
      });
    })
    .then(function (subscription) {
      console.log(
        JSON.stringify({
          subscription: subscription,
        })
      );
    })
    .catch((err) => console.error(err));
}

function urlBase64ToUint8Array(base64String) {
  const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
  const base64 = (base64String + padding)
    .replace(/\-/g, "+")
    .replace(/_/g, "/");
  const rawData = window.atob(base64);
  return Uint8Array.from([...rawData].map((char) => char.charCodeAt(0)));
}

if ("serviceWorker" in navigator) {
  navigator.serviceWorker.register("service-worker.js");
  navigator.serviceWorker.ready
    .then(function (registration) {
      return registration.pushManager.getSubscription();
    })
    .then(function (subscription) {
      if (!subscription) {
        subscribe();
      } else {
        console.log(
          JSON.stringify({
            subscription: subscription,
          })
        );
      }
    });
}
