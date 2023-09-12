document.addEventListener("DOMContentLoaded", function() {
    getIpInfo();
});
function getIpInfo() {
    fetch('http://49.7.114.49:5888/ip')
        .then(response => response.json())
        .then(data => {
            console.log('IP信息:', data);
            // 将获取的IP信息对应到HTML元素
            document.getElementById('country').textContent = data.country;
            document.getElementById('province').textContent = data.region;
            document.getElementById('city').textContent = data.city;
        })
        .catch(error => {
            console.error('获取IP信息失败:', error);
        });
}
function sendSlowRequest() {
    window.open('http://49.7.114.49:5888/slow');
}