<template>
    <div id="cube" ref="stage"></div>
</template>

<script>
import { ref, onMounted, onUnmounted, toRefs } from 'vue';
import * as THREE from 'three';

export default {
    name: 'Cube',
    props: {
        qua: Array,
    },
    setup(props) {
        const { qua } = toRefs(props);
        const stage = ref(null);
        const renderer = ref(null);
        // === scene ===
        const scene = new THREE.Scene();
        // === camera ===
        const camera = new THREE.PerspectiveCamera(30, 1, 0.1, 100);
        scene.add(camera);
        camera.position.z = 5;
        // === light ===
        const light = new THREE.HemisphereLight(0xFFFFFF, 0xFFFFFF, 20.0, new THREE.Vector3(1, 1, 1));
        const directionalLight = new THREE.DirectionalLight(0xFFFFFF, 1);
        directionalLight.position.set(5, 5, 5);
        scene.add(light, directionalLight);


        // === Grid ===
        const grid = new THREE.GridHelper(6, 10, 0x888888, 0x888888);
        grid.position.y = -0.6;
        scene.add(grid);
        // === model ===
        const geometry = new THREE.BoxGeometry(1, 0.5, 0.5);
        const material = [
            new THREE.MeshStandardMaterial({ color: 0x222222, metalness: 0, roughness: 1 }),
            new THREE.MeshStandardMaterial({ color: 0xF4D06F }),
            new THREE.MeshStandardMaterial({ color: 0xFF8811 }),
            new THREE.MeshStandardMaterial({ color: 0x9DD9D2 }),
            new THREE.MeshStandardMaterial({ color: 0x233232 }),
            new THREE.MeshStandardMaterial({ color: 0x235000 }),
        ];
        const cube = new THREE.Mesh(geometry, material);
        const geometry2 = new THREE.SphereGeometry(3, 30, 30);
        const material2 = new THREE.MeshBasicMaterial({ color: 0xffff00, transparent: true, opacity: 0.5 });
        const sphere = new THREE.Mesh(geometry2, material2);
        const lineGeometry = new THREE.BufferGeometry().setFromPoints([
            new THREE.Vector3(-1, 0, 0),
            new THREE.Vector3(10, 0, 0), // You can adjust the length and direction of the line
        ]);

        const lineMaterial = new THREE.LineBasicMaterial({ color: 0x00ff00 });
        const line = new THREE.Line(lineGeometry, lineMaterial);

        sphere.position.set(0, 0, 0);
        cube.position.set(0, 0, 0);
        line.position.set(0, 0, 0);
        scene.add(cube, sphere, line);

        camera.position.set(6, 6, 10);
        camera.lookAt(0, 0, 0);

        function animate() {
            requestAnimationFrame(animate);
            quaternion();
            if (renderer.value && scene && camera) {
                renderer.value.render(scene, camera);
            }
        }

        function quaternion() {
            if (typeof qua.value[0] === 'undefined') {
                console.log('quaternion is undefined');
                return;
            }
            const q = new THREE.Quaternion(qua.value[0], qua.value[1], qua.value[2], qua.value[3]);
            cube.quaternion.copy(q);
            line.quaternion.copy(q);
            // Get the world position of the line
            const lineWorldPosition = new THREE.Vector3();
            line.getWorldPosition(lineWorldPosition);

            // Calculate the direction vector of the line
            const lineDirection = new THREE.Vector3(1, 0, 0); // Assuming the line points in the positive x-direction
            lineDirection.applyQuaternion(q);

            // Calculate the position of the point along the line at a distance of 3 units
            const distance = 3;
            const pointPosition = new THREE.Vector3().copy(lineDirection).multiplyScalar(distance);

            // Create a point (sphere) at the calculated position
            const pointGeometry = new THREE.SphereGeometry(0.05, 16, 16);
            const pointMaterial = new THREE.MeshBasicMaterial({ color: 0xff0000 });
            const point = new THREE.Mesh(pointGeometry, pointMaterial);
            point.position.copy(lineWorldPosition).add(pointPosition);

            // Add the point to the scene
            scene.add(point);
        }

        onMounted(() => {
            renderer.value = new THREE.WebGLRenderer();
            if (stage.value) {
                stage.value.appendChild(renderer.value.domElement);
                renderer.value.setSize(stage.value.offsetWidth / 1.25, 500);
                camera.aspect = (stage.value.offsetWidth / 1.25) / 500;
                camera.updateProjectionMatrix();
                animate();
            }
        });

        onUnmounted(() => {
            // Clean up resources when the component is destroyed.
            // For example, stop animation loops, close connections, etc.
        });

        return { stage, quaternion };
    },
};
</script>
