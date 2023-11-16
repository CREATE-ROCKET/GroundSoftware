import os
import analyze
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation

def process_files(directory="."):
    # Get the absolute path of the directory
    directory = os.path.abspath(directory)

    # Create a list to store the error rates for each file
    all_error_rates_50 = []
    all_total_error_rates = []
    all_files = []

    # Iterate over files in the directory and its subdirectories
    for root, dirs, files in os.walk(directory):
        for file in files:
            # Check if the file is a raw_*.txt file
            if file.startswith("quat_") and file.endswith(".txt"):
                file_path = os.path.join(root, file)
                print(f"Processing file: {file_path}")

                # Calculate error rates for the current file
                error_rates_50, total_error_rate = analyze.calculate_error_rates(file_path)

                # Print error rates for the current file
                print("50行ごとの誤り率:")
                for i, rate in enumerate(error_rates_50):
                    print(f"ブロック {i+1}: {rate * 100:.2f}%")

                print(f"\n全行の誤り率: {total_error_rate * 100:.2f}%")

                # Store error rates for later use
                if total_error_rate != total_error_rate:  # Check if total_error_rate is NaN
                    print("Error: NaN")
                    continue
                
                all_error_rates_50.append(error_rates_50)
                all_total_error_rates.append(total_error_rate)
                all_files.append(file_path)

    # Sort files and error rates in dictionary order
    all_files, all_error_rates_50, all_total_error_rates = zip(*sorted(zip(all_files, all_error_rates_50, all_total_error_rates)))

    # Create an animation to display the error rates over time
    fig, ax = plt.subplots()
    ax.set_ylim(-0.1, 1.1)  # y軸の範囲を設定
    animation = FuncAnimation(fig, animate, frames=len(all_error_rates_50),
                              fargs=(all_error_rates_50, all_total_error_rates, all_files), interval=1)

    # Save the animation as a GIF
    animation_filename = "error_rates_animation.gif"
    animation.save(animation_filename, writer='imagemagick', fps=1)
    print(f"Animation saved as {animation_filename}")

def animate(frame, all_error_rates_50, all_total_error_rates, files):
    # Plot the error rates for the current frame
    error_rates_50 = all_error_rates_50[frame]
    total_error_rate = all_total_error_rates[frame]

    plt.clf()
    plt.plot(error_rates_50, marker='o', label='Error Rate per 50 Rows')
    plt.axhline(y=total_error_rate, color='r', linestyle='-', label='Total Error Rate')

    plt.xlabel('Block Number (Every 50 Rows)')
    plt.ylabel('Error Rate')
    plt.ylim(-0.1,1.1)
    plt.title(f'Error Rates Comparison - File: {files[frame]}')
    plt.legend()

if __name__ == "__main__":
    # Set the directory to the current directory
    current_directory = os.path.dirname(os.path.realpath(__file__))
    process_files(current_directory)
